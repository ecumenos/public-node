package handlers

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

//go:generate mockery --name=Handlers
var _ Handlers = (*handlers)(nil)

// Module is fx module.
var Module = fx.Options(
	fx.Provide(func(log *zap.SugaredLogger) Handlers {
		return &handlers{
			logger: log,
		}
	}),
)

// Handlers is interface for endpoint's handlers.
type Handlers interface {
	GetPing(rw http.ResponseWriter, r *http.Request)
}

type handlers struct {
	logger *zap.SugaredLogger
}
