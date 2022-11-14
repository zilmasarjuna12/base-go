package response

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Message string `json:"message"`
}
