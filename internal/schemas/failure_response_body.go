package schemas

// FailReason defines model for FailReason.
type FailReason struct {
	Data *map[string]interface{} `json:"data"`

	// Description Description is a description of the fail.
	Description string `json:"description"`

	// Status Status is HTTP code.
	Status int `json:"status"`
}

// FailureResponseBody defines model for FailureResponseBody.
type FailureResponseBody struct {
	// Code Code is a numeric code corresponding to the error.
	Code *float32    `json:"code,omitempty"`
	Data *FailReason `json:"data,omitempty"`

	// Message Message is a meaningful, end-user-readable message. It explains what went wrong.
	Message  *string        `json:"message,omitempty"`
	Metadata Metadata       `json:"metadata"`
	Status   ResponseStatus `json:"status"`
}
