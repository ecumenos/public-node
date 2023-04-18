package schemas

// GetPingResponseBody defines model for GetPingResponseBody.
type GetPingResponseBody struct {
	Data     GetPingResponseBodyData `json:"data"`
	Metadata Metadata                `json:"metadata"`
	Status   ResponseStatus          `json:"status"`
}

// GetPingResponseBodyData defines model for GetPingResponseBodyData.
type GetPingResponseBodyData struct {
	Pong bool `json:"pong"`
}
