package handlers

import (
	"net/http"

	"github.com/ecumenos/public-node/httpserver/utilities"
	"github.com/ecumenos/public-node/internal/schemas"
)

// GetPing is handler for GET/ping endpoint.
func (h *handlers) GetPing(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	metadata, err := utilities.GetMetadata(ctx)
	if err != nil {
		utilities.WriteError(ctx, rw, err, h.logger)
		return
	}
	body := schemas.GetPingResponseBody{
		Data: schemas.GetPingResponseBodyData{
			Pong: true,
		},
		Metadata: metadata,
		Status:   schemas.ResponseStatusSuccess,
	}
	utilities.WriteSuccess(rw, h.logger, body)
}
