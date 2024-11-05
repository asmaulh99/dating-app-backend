package errors

import (
	"fmt"
	"net/http"
)

type ErrorType string

const (
	// 400 errors
	ErrBadRequest           ErrorType = "bad_request"          // 400
	ErrUnauthorized         ErrorType = "unauthorized"         // 401
	ErrForbiddenRequest     ErrorType = "forbidden_request"    // 403
	ErrNotFound             ErrorType = "not_found"            // 404
	ErrMethodNotAllowed     ErrorType = "method_not_allowed"   // 405
	ErrUnproccessableEntity ErrorType = "unprocessable_entity" // 422

	ErrInternal ErrorType = "internal" // 500
)

var errorTypeHTTPStatusCodeMap = map[ErrorType]int{
	ErrBadRequest:           http.StatusBadRequest,
	ErrUnauthorized:         http.StatusUnauthorized,
	ErrForbiddenRequest:     http.StatusForbidden,
	ErrNotFound:             http.StatusNotFound,
	ErrMethodNotAllowed:     http.StatusMethodNotAllowed,
	ErrUnproccessableEntity: http.StatusUnprocessableEntity,
	ErrInternal:             http.StatusInternalServerError,
}

type DynamicError struct {
	StatusCode int
	Type       ErrorType
	ErrDetail  string
}

func NewError(errType ErrorType, errDetail string) error {
	return &DynamicError{
		StatusCode: errorTypeHTTPStatusCodeMap[errType],
		Type:       errType,
		ErrDetail:  errDetail,
	}
}

func (e *DynamicError) Error() string {
	return fmt.Sprintf("%s error. %v", e.Type, e.ErrDetail)
}
