package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"wb-tech-l0/internal/logger"
)

// SA1029: should not use built-in type string as key for value; define your own type to avoid collisions
type contextKey string

// key for setting request id in request context
const requestIDKey contextKey = "request_id"

// LoggingMiddleware logs every http request before handlers.
// It sets requestIDKey in request context
func LoggingMiddleware(log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// getting request id
			reqID := r.Header.Get("X-Request-ID")
			if reqID == "" {
				reqID = uuid.NewString()
			}

			// setting to request context
			ctx := context.WithValue(r.Context(), requestIDKey, reqID)

			// checking processing time
			start := time.Now()

			log.Debug("Got HTTP request",
				logger.Field("request_id", reqID),
				logger.Field("method", r.Method),
				logger.Field("path", r.URL.Path),
				logger.Field("remote_addr", r.RemoteAddr),
				logger.Field("user_agent", r.UserAgent()),
			)

			// handling request with context with its id
			next.ServeHTTP(w, r.WithContext(ctx))

			log.Debug("Request processed",
				logger.Field("request_id", reqID),
				logger.Field("duration", time.Since(start).String()),
			)
		})
	}
}

// GetRequestID gets requestIDKey from context
func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return ""
}
