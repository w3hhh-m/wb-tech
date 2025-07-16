package serverhandlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"wb-tech-l0/internal/cache"
	"wb-tech-l0/internal/storage"
)

// GetOrderHandler returns a handler function for getting orders
// It works with cache and storage
func GetOrderHandler(cache cache.Cache, store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// checking method
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// getting uid
		uid := strings.TrimPrefix(r.URL.Path, "/order/")
		if uid == "" {
			http.Error(w, "missing order uid", http.StatusBadRequest)
			return
		}

		// TODO: from cache

		// getting order
		order, err := store.GetOrder(r.Context(), uid)
		if err != nil {
			switch {
			case errors.Is(err, storage.ErrNotFound):
				http.Error(w, "order not found", http.StatusNotFound)
			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}

		// sending response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order) // nolint: errcheck
	}
}
