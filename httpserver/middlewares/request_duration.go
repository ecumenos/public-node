package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ecumenos/golang-toolkit/contexttools"
)

// RequestDurationMiddleware is middleware type for setting start request timestamp value into context.
type RequestDurationMiddleware func(next http.Handler) http.Handler

// NewRequestDurationMiddleware is constructor for RequestDurationMiddleware.
func NewRequestDurationMiddleware() RequestDurationMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			ctx := contexttools.SetValue(r.Context(), contexttools.StartRequestTimestampKey, fmt.Sprint(time.Now().Unix()))

			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
