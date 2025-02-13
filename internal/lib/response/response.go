package response

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

func Error(msg string, status int) Response {
	return Response{
		Status: status,
		Error:  msg,
	}
}
