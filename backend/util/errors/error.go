package errors

import "net/http"

type Error interface {
	Message() string
	Code() int
}

type errorDetail struct {
	msg string
	code int
}

func (d *errorDetail) Message() string {
	return d.msg
}

func (d *errorDetail) Code() int {
	return d.code
}

func NewBadRequestError() Error {
	return &errorDetail{
		msg: http.StatusText(http.StatusBadRequest),
		code: http.StatusBadRequest,
	}
}

func NewInternalServerError() Error {
	return &errorDetail{
		msg: http.StatusText(http.StatusInternalServerError),
		code: http.StatusInternalServerError,
	}
}

func NewNotFoundError() Error {
	return &errorDetail{
		msg: http.StatusText(http.StatusNotFound),
		code: http.StatusNotFound,
	}
}

func NewUnauthorizedError() Error {
	return &errorDetail{
		msg: http.StatusText(http.StatusUnauthorized),
		code: http.StatusUnauthorized,
	}
}

func NewMethodNotAllowedError() Error {
	return &errorDetail{
		msg: http.StatusText(http.StatusMethodNotAllowed),
		code: http.StatusMethodNotAllowed,
	}
}

func NewBadGatewayError() Error {
	return &errorDetail{
		msg: http.StatusText(http.StatusBadGateway),
		code: http.StatusBadGateway,
	}
}