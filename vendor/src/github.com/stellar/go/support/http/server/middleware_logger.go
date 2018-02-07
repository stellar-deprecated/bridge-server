package server

import (
	"math/big"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/stellar/go/support/log"
)

func loggerMiddleware(requestIDKey interface{}, next http.Handler, w http.ResponseWriter, r *http.Request) {
	requestLog := log.WithFields(log.F{
		"request_id": r.Context().Value(requestIDKey),
		"method":     r.Method,
		"uri":        r.RequestURI,
		"ip":         r.RemoteAddr,
	})

	requestLog.Info("HTTP request")

	ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
	requestStartTime := time.Now()

	next.ServeHTTP(ww, r)

	duration := big.NewRat(
		time.Since(requestStartTime).Nanoseconds(),
		int64(time.Second),
	)

	requestLog.WithFields(log.F{
		"status":         ww.Status(),
		"response_bytes": ww.BytesWritten(),
		"duration":       duration.FloatString(8),
	}).Info("HTTP response")
}

// BindLoggerMiddleware returns a LoggerMiddleware bound to the passed in requestIdKey, thereby decoupling it from middleware.RequestID
func BindLoggerMiddleware(requestIDKey interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loggerMiddleware(requestIDKey, next, w, r)
		})
	}
}
