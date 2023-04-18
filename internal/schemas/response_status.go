package schemas

// ResponseStatus defines model for ResponseStatus.
type ResponseStatus string

// Defines values for ResponseStatus.
const (
	ResponseStatusError   ResponseStatus = "error"
	ResponseStatusFail    ResponseStatus = "fail"
	ResponseStatusSuccess ResponseStatus = "success"
)
