package module

import (
	"github.com/ecumenos/public-node/clientsledger"
	"github.com/ecumenos/public-node/config"
	"github.com/ecumenos/public-node/httpserver"
	"github.com/ecumenos/public-node/internal/env"
	"github.com/ecumenos/public-node/internal/logger"
	"github.com/ecumenos/public-node/internal/types"
	"github.com/ecumenos/public-node/nodesledger"
	"github.com/ecumenos/public-node/secretmanager"
	"go.uber.org/fx"
)

// NewModule is constructor for general fx module that take service name.
func NewModule(serviceName string) fx.Option {
	return fx.Options(
		fx.Supply(types.SericeVersion(config.ServiceVersion)),
		fx.Supply(types.ServiceName(serviceName)),
		logger.Module,
		env.Module(&config.Config{}, false),
		fx.Provide(func(c env.Config) config.Config {
			return *c.(*config.Config)
		}),
		httpserver.Module,
		clientsledger.Module,
		nodesledger.Module,
		secretmanager.Module,
	)
}
