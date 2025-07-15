package serverhandlers

import (
	"net/http"
	"wb-tech-l0/internal/cache"
	"wb-tech-l0/internal/storage"
)

// GetOrderHandler returns a handler function for getting orders
// It works with cache and storage
func GetOrderHandler(cache cache.Cache, storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement me
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Not implemented")) // nolint: errcheck
	}
}
