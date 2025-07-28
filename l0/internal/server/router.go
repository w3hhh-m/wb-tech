package server

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"wb-tech-l0/internal/cache"
	"wb-tech-l0/internal/logger"
	serverHandlers "wb-tech-l0/internal/server/handlers"
	"wb-tech-l0/internal/server/middlewares"
	"wb-tech-l0/internal/storage"
)

// NewRouter creates and returns a new HTTP router with all handlers registered
func NewRouter(log logger.Logger, cache cache.Cache, storage storage.Storage) http.Handler {
	mux := http.NewServeMux()
	// register GetOrder handler
	mux.HandleFunc("/api/order/", serverHandlers.GetOrderHandler(log, cache, storage))
	// Swagger docs handler
	mux.HandleFunc("/api/docs/", httpSwagger.WrapHandler)
	// adding logger middleware
	return middlewares.LoggingMiddleware(log)(mux)
}
