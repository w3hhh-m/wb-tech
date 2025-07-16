package server

import (
	"net/http"
	"wb-tech-l0/internal/cache"
	serverHandlers "wb-tech-l0/internal/server/handlers"
	"wb-tech-l0/internal/storage"
)

// NewRouter creates and returns a new HTTP router with all handlers registered
func NewRouter(cache cache.Cache, storage storage.Storage) http.Handler {
	mux := http.NewServeMux()
	// register GetOrder handler
	mux.HandleFunc("/order/", serverHandlers.GetOrderHandler(cache, storage))
	return mux
}
