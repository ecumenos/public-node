package middlewares

import (
	"net/http"

	"github.com/ecumenos/golang-toolkit/contexttools"
	"github.com/ecumenos/golang-toolkit/randomtools"
)

// CorrelationIDMiddleware is middleware type for setting correlation_id into context.
type CorrelationIDMiddleware func(next http.Handler) http.Handler

// NewCorrelationIDMiddleware is constructor for CorrelationIDMiddleware.
func NewCorrelationIDMiddleware() CorrelationIDMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			var correlationID string
			if cID := r.Header.Get("X-Correlation-ID"); cID != "" {
				correlationID = cID
			} else {
				correlationID = randomtools.GetUUIDString()
			}
			ctx := contexttools.SetValue(r.Context(), contexttools.CorrelationIDKey, correlationID)

			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
