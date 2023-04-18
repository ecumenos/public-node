package utilities

import (
	"net/http"

	"go.uber.org/zap"
)

// WriteSuccess writes HTTP response with embeded 200 OK status code.
func WriteSuccess(rw http.ResponseWriter, log *zap.SugaredLogger, body interface{}) {
	WriteResponse(log, body, http.StatusOK, rw)
}
