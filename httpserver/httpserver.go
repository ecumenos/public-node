package httpserver

import (
	"context"
	"net/http"

	"github.com/ecumenos/public-node/httpserver/handlers"
	"github.com/ecumenos/public-node/httpserver/middlewares"
	"github.com/ecumenos/public-node/internal/types"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module is fx module that invoke starting HTTP server.
var Module = fx.Options(
	handlers.Module,
	fx.Provide(
		middlewares.NewCorrelationIDMiddleware,
		middlewares.NewIPAddressMiddleware,
		middlewares.NewRequestDurationMiddleware,
		NewRouter,
	),
	fx.Invoke(Start),
)

const (
	timeoutMessage = `{"status": "error", "message": "server timeout", "code": 503}`
)

// Params is input parameters for starting HTTP server.
type Params struct {
	fx.In
	Lifecycle   fx.Lifecycle
	ServiceName types.ServiceName
	Config      Config
	Log         *zap.Logger
	Router      *mux.Router
	Shutdowner  fx.Shutdowner
}

// Start is function for starting HTTP server.
func Start(params Params) error {
	var handler http.Handler = params.Router
	timeoutHandler := http.TimeoutHandler(handler, params.Config.HandlerTimeout, timeoutMessage)
	httpServer := &http.Server{
		Addr:         params.Config.Address,
		WriteTimeout: params.Config.WriteTimeout,
		ReadTimeout:  params.Config.ReadTimeout,
		IdleTimeout:  params.Config.IdleTimeout,
		Handler:      timeoutHandler,
	}
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				params.Log.Info("http server is starting...")
				shutdownStatus := httpServer.ListenAndServe()
				params.Log.Info("http server shutdown status", zap.Any("status", shutdownStatus))
				if err := params.Shutdowner.Shutdown(); err != nil {
					params.Log.Error("shutdown http server error", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Log.Info("http server is shutting down...")
			if err := httpServer.Shutdown(ctx); err != nil {
				params.Log.Error("shutting down http server error", zap.Error(err))
				return err
			}

			return nil
		},
	})

	return nil
}
