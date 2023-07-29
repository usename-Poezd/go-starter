package responses

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	Errors []interface{} `json:"errors"`
}

type Response struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
