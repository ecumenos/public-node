package utilities

import (
	"context"
	"net/http"

	"github.com/ecumenos/golang-toolkit/customerror"
	"github.com/ecumenos/golang-toolkit/ptrtools"
	"github.com/ecumenos/public-node/internal/schemas"
	"go.uber.org/zap"
)

// WriteError writes error in HTTP response.
func WriteError(ctx context.Context, rw http.ResponseWriter, err error, log *zap.SugaredLogger) {
	if err == nil {
		return
	}

	customerror.Log(log, err)

	ce := customerror.Cast(err)
	switch ce.Class {
	case customerror.ErrorClass:
		writeErrorResp(ctx, rw, ce.Err, log)
	case customerror.FailureClass:
		writeFailureResp(ctx, rw, ce.Fail, log)
	}
}

func writeErrorResp(ctx context.Context, rw http.ResponseWriter, err customerror.Error, log *zap.SugaredLogger) {
	metadata, merr := GetMetadata(ctx)
	if merr != nil {
		log.Error("failed to get metadata", zap.Error(merr))
		return
	}

	body := schemas.ErrorResponseBody{
		Status:   "error",
		Code:     ptrtools.ValueToPtr(float32(err.Code.UInt32())),
		Message:  err.Message,
		Metadata: metadata,
	}
	WriteResponse(log, body, http.StatusInternalServerError, rw)
}

func writeFailureResp(ctx context.Context, rw http.ResponseWriter, err customerror.Failure, log *zap.SugaredLogger) {
	var data *map[string]interface{}
	if err.Data != nil {
		data = &err.Data
	}

	metadata, merr := GetMetadata(ctx)
	if merr != nil {
		log.Error("failed to get metadata", zap.Error(merr))
		return
	}

	body := schemas.FailureResponseBody{
		Status: schemas.ResponseStatusFail,
		Code:   ptrtools.ValueToPtr(float32(err.Code.UInt32())),
		Data: ptrtools.ValueToPtr(schemas.FailReason{
			Data:        data,
			Description: err.Description,
			Status:      err.StatusCode,
		}),
		Message:  ptrtools.ValueToPtr(err.Message),
		Metadata: metadata,
	}
	WriteResponse(log, body, http.StatusBadRequest, rw)
}
