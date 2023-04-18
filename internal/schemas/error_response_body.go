package schemas

// ErrorResponseBody defines model for ErrorResponseBody.
type ErrorResponseBody struct {
	// Code Code is a numeric code corresponding to the error.
	Code *float32 `json:"code,omitempty"`

	// Message Message is a meaningful, end-user-readable message. It explains what went wrong.
	Message  string         `json:"message"`
	Metadata Metadata       `json:"metadata"`
	Status   ResponseStatus `json:"status"`
}
