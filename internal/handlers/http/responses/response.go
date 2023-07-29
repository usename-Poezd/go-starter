package responses

type ErrorResponse struct {
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors,omitempty"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}
