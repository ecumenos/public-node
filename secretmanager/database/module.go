package database

import "go.uber.org/fx"

// Module is fx module.
var Module = fx.Options(
	fx.Provide(NewConnection),
)
