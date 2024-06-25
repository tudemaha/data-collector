package dto

import "net/http"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (r *Response) OK() {
	r.Status = http.StatusOK
	r.Message = "success"
}

func (r *Response) BadRequest() {
	r.Status = http.StatusBadRequest
}

func (r *Response) MethodNotAllowed() {
	r.Status = http.StatusMethodNotAllowed
	r.Message = "method not allowed"
}

func (r *Response) InternalServerError() {
	r.Status = http.StatusInternalServerError
}
