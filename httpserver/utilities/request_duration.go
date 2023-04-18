package utilities

import (
	"context"
	"time"

	"github.com/ecumenos/golang-toolkit/contexttools"
	"github.com/ecumenos/golang-toolkit/customerror"
	"github.com/ecumenos/golang-toolkit/numbertools"
)

// GetRequestDuration returns duration of request based on context.
func GetRequestDuration(ctx context.Context) (int, error) {
	str := contexttools.GetValue(ctx, contexttools.StartRequestTimestampKey)
	if str == "" {
		return 0, nil
	}
	startDateUnix, err := numbertools.StringToInt64(str)
	if err != nil {
		return 0, customerror.NewError(err, "can not parse start request duration datetime", customerror.DefaultErrorCode)
	}

	return int(time.Now().Unix() - startDateUnix), nil
}
