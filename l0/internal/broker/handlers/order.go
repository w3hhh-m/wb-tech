package brokerhandlers

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"

	"wb-tech-l0/internal/broker"
	"wb-tech-l0/internal/logger"
	"wb-tech-l0/internal/models"
	"wb-tech-l0/internal/storage"
)

// OrdersHandler returns a handler function for broker.Subscribe for handling orders messages.
// handler must return error if something is wrong with the message handling.
// on error, broker will NOT commit message and there could be retries.
func OrdersHandler(log logger.Logger, store storage.Storage, validate *validator.Validate) func(message *broker.Message) error {
	return func(message *broker.Message) error {
		// add message key to log
		log := log.With(logger.Field("message_key", string(message.Key)))

		var order models.Order
		// parsing message value in order struct
		if err := json.Unmarshal(message.Value, &order); err != nil {
			log.Debug("Invalid JSON message. Handler skipping message", logger.Error(err))
			// returning nil to commit message in Subscribe
			return nil
		}

		// validating
		if err := validate.Struct(order); err != nil {
			log.Debug("Invalid order schema. Handler skipping message", logger.Error(err))
			// returning nil to commit message in Subscribe
			return nil
		}

		// adding order uid to logger for chaining storage logs with handler logs
		log = log.With(logger.Field("order_uid", order.OrderUID))

		// saving message
		err := store.SaveOrder(&order)
		if err != nil {
			log.Warn("Failed to save order", logger.Error(err))
			if errors.Is(err, storage.ErrUniqueViolation) {
				log.Warn("Skipping order because it already exists")
				// returning nil to commit message in Subscribe because of invalid data
				return nil
			}
			// returning error to NOT commit message in broker
			return err
		}

		return nil
	}
}
