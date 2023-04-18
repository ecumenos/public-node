package logger

import (
	"time"

	toolkitlogger "github.com/ecumenos/golang-toolkit/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newLogger sets timestamp key to "time" and encoding to ISO8601 with millisecond precision
func newLogger(serviceName string, zapConfig zap.Config) (*zap.Logger, error) {
	zapConfig.EncoderConfig.TimeKey = "time"
	zapConfig.EncoderConfig.EncodeTime = utcISO8601TimeEncoder

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	logger = logger.WithOptions(zap.Fields(zap.String(toolkitlogger.ServiceLoggerField, serviceName)))
	return logger, nil
}

// utcISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string
// with millisecond precision in UTC
func utcISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z0700"))
}

// NewProductionLogger returns a zap.NewProductionConfig logger https://godoc.org/go.uber.org/zap#NewProductionConfig
// with the service name and ISO8601 utc timestamp
//
// Production logger uses default zap production sampling rules. For logs with the same message and
// level the first 100 logs per second are logged, then the following 99 logs are silently dropped
// with only the next 100th log emitted.
func NewProductionLogger(serviceName string) (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	return newLogger(serviceName, zapConfig)
}

// NewDevelopmentLogger is the same as NewProductionLogger, but with debug level and no sampling.
func NewDevelopmentLogger(serviceName string) (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zapConfig.Sampling = nil
	return newLogger(serviceName, zapConfig)
}
