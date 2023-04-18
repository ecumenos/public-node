package config

import (
	clientsledgerConfig "github.com/ecumenos/public-node/clientsledger/config"
	"github.com/ecumenos/public-node/httpserver"
	"github.com/ecumenos/public-node/internal/logger"
	nodesledgerConfig "github.com/ecumenos/public-node/nodesledger/config"
	secretmanagerConfig "github.com/ecumenos/public-node/secretmanager/config"
	"go.uber.org/fx"
)

// Config is common configuration for all application.
type Config struct {
	fx.Out
	Log                 logger.Config
	HTTPServer          httpserver.Config
	ClientsLedger       clientsledgerConfig.Config
	NodesLedger         nodesledgerConfig.Config
	SecretManagerLedger secretmanagerConfig.Config
}
