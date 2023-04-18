package utilities

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// WriteResponse writes HTTP response by response body & status code.
func WriteResponse(log *zap.SugaredLogger, body interface{}, statusCode int, rw http.ResponseWriter) {
	b, err := json.Marshal(body)
	if err != nil {
		log.Error("failed to marshal response", zap.Error(err))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	_, err = rw.Write(b)
	if err != nil {
		log.Error("failed to write response", zap.Error(err))
	}
}
