package serverhandlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"wb-tech-l0/internal/cache"
	"wb-tech-l0/internal/logger"
	"wb-tech-l0/internal/models"
	"wb-tech-l0/internal/server/middlewares"
	"wb-tech-l0/internal/storage"
)

// GetOrderHandler godoc
//
//	@Summary		Получить заказ по UID
//	@Description	Возвращает заказ по его уникальному идентификатору
//	@Tags			order
//	@Param			order_uid	path		string	true	"UID заказа"
//	@Success		200			{object}	models.Order
//	@Failure		400			{string}	string	"missing order uid"
//	@Failure		404			{string}	string	"order not found"
//	@Failure		405			{string}	string	"method not allowed"
//	@Failure		500			{string}	string	"internal server error"
//	@Router			/order/{order_uid} [get]
func GetOrderHandler(log logger.Logger, cache cache.Cache, store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// getting request id
		requestID := middlewares.GetRequestID(r.Context())
		log = log.With(logger.Field("request_id", requestID))

		// checking method
		if r.Method != http.MethodGet {
			log.Debug("Request method is not allowed")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// getting uid
		uid := strings.TrimPrefix(r.URL.Path, "/order/")
		if uid == "" {
			log.Debug("Request /{order_uid} path is empty")
			http.Error(w, "missing order uid", http.StatusBadRequest)
			return
		}

		// try to get from cache first
		if cached, found := cache.GetOrder(uid); found {
			var ok bool
			// check if it iss order
			if cached, ok = cached.(*models.Order); !ok {
				log.Debug("Requested item in cache is not order", logger.Field("uid", uid))
				http.Error(w, "order not found", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cached) // nolint: errcheck
			log.Debug("Successfully sent order response from cache")
			return
		}

		// getting order
		order, err := store.GetOrder(r.Context(), uid)
		if err != nil {
			switch {
			case errors.Is(err, storage.ErrNotFound):
				http.Error(w, "order not found", http.StatusNotFound)
			default:
				log.Warn("Failed to get order", logger.Error(err))
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}

		// save to cache for future requests
		cache.SaveOrder(uid, order)

		// sending response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order) // nolint: errcheck

		log.Debug("Successfully sent order response")
	}
}
