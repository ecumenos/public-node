package utilities

import (
	"context"
	"time"

	"github.com/ecumenos/golang-toolkit/contexttools"
	"github.com/ecumenos/golang-toolkit/timetools"
	"github.com/ecumenos/public-node/internal/schemas"
)

// GetMetadata returns Metadata based on context.
func GetMetadata(ctx context.Context) (schemas.Metadata, error) {
	duration, err := GetRequestDuration(ctx)
	if err != nil {
		return schemas.Metadata{}, err
	}

	return schemas.Metadata{
		CorrelationID: contexttools.GetValue(ctx, contexttools.CorrelationIDKey),
		Timestamp:     timetools.TimeToString(time.Now()),
		Duration:      duration,
		Version:       contexttools.GetValue(ctx, contexttools.ServiceVersionKey),
	}, nil
}
