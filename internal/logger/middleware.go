package logger

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestID := middleware.GetReqID(r.Context())
		if requestID == "" {
			requestID = middleware.RequestIDHeader
		}

		reqLogger := WithRequestID(requestID)

		ctx := NewContext(r.Context(), reqLogger)
		r = r.WithContext(ctx)

		reqLogger.Debug("Request started",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent())

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		duration := time.Since(start)
		reqLogger.Debug("Request completed",
			"status", ww.Status(),
			"bytes", ww.BytesWritten(),
			"duration_ms", float64(duration.Microseconds())/1000)
	})
}
