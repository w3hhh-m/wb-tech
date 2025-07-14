package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wb-tech-l0/internal/logger"
	"wb-tech-l0/internal/models"
	"wb-tech-l0/internal/storage"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres is a Storage interface implementation for PostgreSQL
type Postgres struct {
	pool           *pgxpool.Pool
	requestTimeout time.Duration
	retryTimeout   time.Duration
	maxRetries     int

	ctx context.Context
	log logger.Logger
}

// New creates and returns initialized Postgres implementation of Storage interface
func New(ctx context.Context, cfg *Config, log logger.Logger) (*Postgres, error) {
	log.Debug("Creating storage connection")

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)

	dbpoolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("error parsing postgres storage config: %w", err)
	}

	dbpoolCfg.ConnConfig.ConnectTimeout = cfg.ConnectTimeout
	dbpoolCfg.MaxConns = cfg.PoolMaxConns
	dbpoolCfg.MinConns = cfg.PoolMinConns
	dbpoolCfg.MaxConnLifetime = cfg.MaxConnLifetime
	dbpoolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, dbpoolCfg)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres storage: %w", err)
	}

	return &Postgres{
		pool:           pool,
		requestTimeout: cfg.RequestTimeout,
		retryTimeout:   cfg.RetryTimeout,
		maxRetries:     cfg.MaxRetries,
		log:            log,
		ctx:            ctx,
	}, nil
}

// Close closes the Postgres storage connection
func (p *Postgres) Close() error {
	p.pool.Close()
	return nil
}

// SaveOrder takes order and tries to save it max retries times or until success.
// It returns error if after max retires times order still was not saved.
func (p *Postgres) SaveOrder(order *models.Order) error {
	var err error

	// adding order uid to logs for chaining with handler logs
	log := p.log.With(logger.Field("order_uid", order.OrderUID))
	// adding max attempts to logs
	log = log.With(logger.Field("max_attempts", p.maxRetries))

	// saving with max retries
	for attempt := 1; attempt <= p.maxRetries; attempt++ {

		log.Debug("Attempting to save order", logger.Field("attempt", attempt))

		// creating context for this retry with request timeout
		ctx, cancel := context.WithTimeout(p.ctx, p.requestTimeout)

		// using function, to defer context cancel and rollback on error
		func() {
			defer cancel()

			// begin the transaction
			tx, txErr := p.pool.BeginTx(ctx, pgx.TxOptions{})
			if txErr != nil {
				err = fmt.Errorf("could not begin tx: %w", txErr)
				return
			}

			defer func() {
				// transaction must be rolled back even if request context is timed out
				// so using application context

				// rollback is safe to call if commit was successful
				err := tx.Rollback(p.ctx)
				if err != nil {
					// pgx.ErrTxClosed is when transaction already closed (on Commit)
					if !errors.Is(err, pgx.ErrTxClosed) {
						log.Warn("Failed to rollback transaction", logger.Field("attempt", attempt), logger.Error(err))
					}
				}
			}()

			// inserting
			err = p.insertOrderTx(ctx, tx, order)
			if err != nil {
				log.Warn("Failed to save order", logger.Field("attempt", attempt), logger.Error(err))
				return
			}

			// commiting transaction
			err = tx.Commit(ctx)
			if err != nil {
				log.Warn("Failed to commit transaction", logger.Field("attempt", attempt), logger.Error(err))
				return
			}
		}()

		if err == nil {
			// if everything was good, return nil error
			return nil
		}

		// if error is about sql
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// if message violates unique constraint, return this error
			if pgErr.Code == pgerrcode.UniqueViolation {
				return storage.ErrUniqueViolation
			}
		}

		// waiting for next try or app context cancellation
		if attempt < p.maxRetries {
			select {
			case <-p.ctx.Done():
				return p.ctx.Err()
			case <-time.After(p.retryTimeout):
				// continue retries
			}
		}
	}

	return fmt.Errorf("save order failed after %d attempts: %w", p.maxRetries, err)
}

// insertOrderTx is a helper method to insert order within a given transaction
// It returns error if something goes wrong. In that case, transaction must be
// rolled back by function that called this method
func (p *Postgres) insertOrderTx(ctx context.Context, tx pgx.Tx, o *models.Order) error {
	// inserting order
	_, err := tx.Exec(ctx, `
		INSERT INTO orders (
			order_uid, track_number, entry, locale, internal_signature,
			customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`,
		o.OrderUID, o.TrackNumber, o.Entry, o.Locale, o.InternalSignature,
		o.CustomerID, o.DeliveryService, o.ShardKey, o.SmID, o.DateCreated, o.OofShard,
	)
	if err != nil {
		return fmt.Errorf("could not insert orders: %w", err)
	}

	// inserting delivery
	_, err = tx.Exec(ctx, `
		INSERT INTO delivery (
			order_uid, name, phone, zip, city, address, region, email
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`,
		o.OrderUID, o.Delivery.Name, o.Delivery.Phone, o.Delivery.Zip,
		o.Delivery.City, o.Delivery.Address, o.Delivery.Region, o.Delivery.Email,
	)
	if err != nil {
		return fmt.Errorf("could not insert delivery: %w", err)
	}

	// inserting payment
	_, err = tx.Exec(ctx, `
		INSERT INTO payment (
			transaction, order_uid, request_id, currency, provider,
			amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`,
		o.Payment.Transaction, o.OrderUID, o.Payment.RequestID,
		o.Payment.Currency, o.Payment.Provider, o.Payment.Amount,
		o.Payment.PaymentDT, o.Payment.Bank, o.Payment.DeliveryCost,
		o.Payment.GoodsTotal, o.Payment.CustomFee,
	)
	if err != nil {
		return fmt.Errorf("could not insert payment: %w", err)
	}

	// inserting items
	for _, item := range o.Items {
		_, err := tx.Exec(ctx, `
			INSERT INTO items (
				order_uid, chrt_id, track_number, price, rid,
				name, sale, size, total_price, nm_id, brand, status
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		`,
			o.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.RID,
			item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status,
		)
		if err != nil {
			return fmt.Errorf("could not insert item: %w", err)
		}
	}

	return nil
}
