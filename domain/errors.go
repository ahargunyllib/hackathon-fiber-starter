package domain

import (
	"errors"
	"net/http"
)

type SerializableError interface {
	Serialize() any
}

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

var ErrNotFound = &RequestError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("something not found"),
}

var ErrNoAPIKey = &RequestError{
	StatusCode: http.StatusUnauthorized,
	Err:        errors.New("no api key provided"),
}

var ErrInvalidAPIKey = &RequestError{
	StatusCode: http.StatusUnauthorized,
	Err:        errors.New("invalid api key"),
}

var ErrUserNotFound = &RequestError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("user not found"),
}

var ErrUserEmailAlreadyExists = &RequestError{
	StatusCode: http.StatusConflict,
	Err:        errors.New("user email already exists"),
}
