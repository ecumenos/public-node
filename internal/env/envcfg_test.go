package env

import (
	"os"
	"testing"

	"github.com/ecumenos/public-node/internal/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/goleak"
)

// TestConfig is structure for test config.
type TestConfig struct {
	Test string
}

func init() {
	// To test the default configuration, we need to set all the
	// required env vars and run the NewConfiguration function to
	// generate a config with all the defaults.
	os.Clearenv()
	_ = os.Setenv("TEST", "testWithoutPrefix")
	_ = os.Setenv("TESTSERVICENAME_TEST", "testWithPrefix")
}

// ConfigInvoke is a fake invoke that will cause the config to be populated.
func ConfigInvoke(_ Config) {}

func Test_EnvModule(t *testing.T) {
	defer goleak.VerifyNone(t)

	cfg := &TestConfig{}
	app := fxtest.New(t,
		fx.Supply(
			types.ServiceName("TestServiceName"),
		),
		Module(cfg, false),
		fx.Invoke(ConfigInvoke),
	)

	app.RequireStart()
	assert.Equal(t, cfg.Test, "testWithoutPrefix")
	app.RequireStop()
}

func Test_EnvServiceNamePrefix(t *testing.T) {
	defer goleak.VerifyNone(t)

	cfg := &TestConfig{}
	app := fxtest.New(t,
		fx.Supply(
			types.ServiceName("TestServiceName"),
		),
		// Don't use module here as go panics when NewConfigOnlyCheck redefines flags when running
		// multiple tests
		fx.Provide(NewEnvConfig(cfg, true)),
		fx.Invoke(ConfigInvoke),
	)

	app.RequireStart()
	assert.Equal(t, cfg.Test, "testWithPrefix")
	app.RequireStop()
}
