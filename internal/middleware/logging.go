package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func Logging(log *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			rw := NewResponseWriter(w)
			next.ServeHTTP(rw, r)

			log.Debug("",
				slog.String("op", "loggingMW"),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("estimate", fmt.Sprintf("%d ms", time.Since(startTime)/1_000_000)),
				slog.String("requestID", middleware.GetReqID(r.Context())),
				slog.Int("statusCode", rw.statusCode),
			)
		})
	}
}
