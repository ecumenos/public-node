package schemas

// Metadata defines model for Metadata.
type Metadata struct {
	// CorrelationId CorrelationID is an ID of request.
	CorrelationID string `json:"correlation_id"`

	// Duration Duration is duration of processing request in milliseconds.
	Duration int `json:"duration"`

	// Timestamp Timestamp is datetime of end of processing request. Format 2006-01-02T15:04:05Z07:00 .
	Timestamp string `json:"timestamp"`

	// Version Version is semver version of a service.
	Version string `json:"version"`
}
