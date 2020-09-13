package dtos

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
}
