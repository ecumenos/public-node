package types

import (
	"go.uber.org/fx"
)

// ServiceName is type of service name.
type ServiceName string

// SericeVersion is type of service version.
type SericeVersion string

// Flag is type for flag.
type Flag struct{}

// Flags is fx input type with list of flags.
type Flags struct {
	fx.In
	Flugs []Flag `group:"flag"`
}
