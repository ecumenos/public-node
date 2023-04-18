package httpserver

import (
	"net/http"

	"github.com/ecumenos/golang-toolkit/contexttools"
	"github.com/ecumenos/golang-toolkit/logger"
	"github.com/ecumenos/public-node/httpserver/handlers"
	"github.com/ecumenos/public-node/httpserver/middlewares"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewRouterParams is input parameters for creation of router for HTTP server.
type NewRouterParams struct {
	fx.In
	Handlers                  handlers.Handlers
	Logger                    *zap.SugaredLogger
	CorrelationIDMiddleware   middlewares.CorrelationIDMiddleware
	IPAddressMiddleware       middlewares.IPAddressMiddleware
	RequestDurationMiddleware middlewares.RequestDurationMiddleware
}

// NewRouter is constrctor of router for HTTP server.
func NewRouter(params NewRouterParams) *mux.Router {
	r := mux.NewRouter()
	r.Use(
		mux.MiddlewareFunc(params.RequestDurationMiddleware),
		mux.MiddlewareFunc(params.CorrelationIDMiddleware),
		mux.MiddlewareFunc(params.IPAddressMiddleware),
	)

	r.HandleFunc("/ping", params.Handlers.GetPing)

	r.Use(
		mux.MiddlewareFunc(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				ctx := r.Context()

				params.Logger.Infow("receiving request",
					zap.String(logger.CorrelationIDLoggerField, contexttools.GetValue(ctx, contexttools.CorrelationIDKey)),
					zap.String("request-ip-address", contexttools.GetValue(ctx, contexttools.IPAddressKey)),
					zap.String("uri", r.RequestURI),
				)
				next.ServeHTTP(rw, r)
			})
		}),
	)

	return r
}
