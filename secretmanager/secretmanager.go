package secretmanager

import (
	"github.com/ecumenos/public-node/secretmanager/database"
	"go.uber.org/fx"
)

//go:generate mockery --name=SecretManager
var _ SecretManager = (*manager)(nil)

// SecretManager is interface for secret manager logic.
type SecretManager interface{}

// Module is fx module.
var Module = fx.Options(
	database.Module,
	fx.Provide(func() SecretManager {
		return &manager{}
	}),
)

type manager struct{}
