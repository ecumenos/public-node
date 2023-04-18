package env

import (
	"flag"
	"fmt"
	"log"

	"github.com/ecumenos/public-node/internal/types"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/fx"
)

// Module bundles important provide/invoke in this package into an fx.Option to allow easy addition to a service.
// If useServicePrefix is set to true all envconfig variables will be prefixed with the service name
var Module = func(cfg interface{}, useServicePrefix bool) fx.Option {
	return fx.Options(
		fx.Provide(NewEnvConfig(cfg, useServicePrefix)),
		fx.Invoke(
			NewConfigOnlyCheck(cfg),
		),
	)
}

// Config is a wrapper of a service specific config struct.
type Config interface{}

// NewEnvConfig provides service the specific config struct from a `envconf` defined struct. If
// useServicePrefix is set to true all env variables will be prefixed with the service name
func NewEnvConfig(cfg interface{}, useServicePrefix bool) func(sn types.ServiceName) (Config, error) {
	return func(sn types.ServiceName) (Config, error) {
		prefix := ""
		if useServicePrefix {
			prefix = string(sn)
		}
		if err := envconfig.Process(prefix, cfg); err != nil {
			return cfg, fmt.Errorf("error: parsing config: %s", err)
		}
		return cfg, nil
	}
}

// NewConfigOnlyCheck will check if we should print out flags/usage only and exit.
func NewConfigOnlyCheck(cfg interface{}) func(sn types.ServiceName, _ types.Flags) {
	return func(sn types.ServiceName, _ types.Flags) {
		flag.Usage = func() {
			fmt.Printf("This daemon is a service which manages %s.\n\nUsage of %s:\n\n%s [flags]\n\n", sn, sn, sn)
			flag.PrintDefaults()
			fmt.Printf("\nConfiguration:\n\n")
			err := envconfig.Usage("", cfg)
			if err != nil {
				fmt.Printf("Unable to dump usage: %s", err)
			}
		}
		var configOnly bool
		flag.BoolVar(&configOnly, "config-only", false, "only show parsed configuration and exit")
		flag.Parse()
		if configOnly {
			err := envconfig.Usage("", cfg)
			if err != nil {
				fmt.Printf("Unable to dump usage: %s", err)
			}
			log.Fatalf("exit due to config only flag")
		}
	}
}
