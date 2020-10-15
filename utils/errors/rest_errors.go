package errors

import "net/http"

//RestErr rest error
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

//NewBadRequestError Bad request error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   message,
	}
}

//NewNotFoundError not found
func NewNotFoundError(message string) *RestErr {
	return &(RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   message,
	})
}
