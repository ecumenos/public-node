package middlewares

import (
	"net/http"

	"github.com/ecumenos/golang-toolkit/contexttools"
	"github.com/ecumenos/golang-toolkit/httptools"
	"go.uber.org/zap"
)

// IPAddressMiddleware is middleware type for setting ip_address into context.
type IPAddressMiddleware func(next http.Handler) http.Handler

// NewIPAddressMiddleware is constructor for IPAddressMiddleware.
func NewIPAddressMiddleware(log *zap.SugaredLogger) IPAddressMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			ip, err := httptools.ExtractIPAddress(r)
			if err != nil {
				log.Errorw("can not extract IP address from request", zap.Error(err))
			}
			ctx := contexttools.SetValue(r.Context(), contexttools.IPAddressKey, ip)

			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
