package errors

import "net/http"

type StatusError struct {
	Code    int
	Message string
}

func New(code int, msg string) *StatusError {
	if "" == msg {
		switch code {
		case http.StatusBadRequest:
			msg = "Missing parameters"
		case http.StatusUnauthorized:
			msg = "Invalid credentials"
		case http.StatusForbidden:
			msg = "Forbidden"
		case http.StatusNotFound:
			msg = "Resource not found"
		case http.StatusMethodNotAllowed:
			msg = "Method not supported by resource"
		case http.StatusUnsupportedMediaType:
			msg = "Content type not supported by resource"
		case http.StatusUnprocessableEntity:
			msg = "Missing required information"
		case http.StatusNotImplemented:
			msg = "Method not implemented"
		default:
			msg = "Internal error"
		}
	}
	return &StatusError{
		Code:    code,
		Message: msg,
	}
}

func NewInternalServerError(msg string) *StatusError {
	return &StatusError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func (err *StatusError) Error() string {
	return err.Message
}
