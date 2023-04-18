package logger

import (
	"context"

	"go.uber.org/fx/fxevent"

	"github.com/ecumenos/public-node/internal/types"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Config defines log related configurations.
type Config struct {
	Production bool `default:"true" desc:"Production vs development logging."`
}

// Module bundles important provide/invoke in this package into a fx.Option.
var Module = fx.Options(
	fx.Provide(
		NewZapLogger,
		ZapSugared,
	),
	// Sets the fx logger to use the zap logger above. This way any logs from fx get bundled with service logs.
	// This additionally tags them as info/error as appropriate making it easy to filter out the fx logs that are
	// not important
	fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: logger}
	}),
)

// NewZapLogger provides a logger using the uber zap library. This function will additionally switch out the global logger
// the logger this function returns
func NewZapLogger(serviceName types.ServiceName, cfg Config, lc fx.Lifecycle) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if cfg.Production {
		logger, err = NewProductionLogger(string(serviceName))
	} else {
		logger, err = NewDevelopmentLogger(string(serviceName))
	}
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)

	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			_ = logger.Sync()
			return nil
		},
	})

	return logger, nil
}

// ZapSugared provides an automatic way to switch from a logger to a sugared logger
func ZapSugared(log *zap.Logger) *zap.SugaredLogger {
	return log.Sugar()
}
