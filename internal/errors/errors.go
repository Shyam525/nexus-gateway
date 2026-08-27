// package errors

// // TODO: Implement errors



package errors

type Code string

const (
	InvalidRequest Code = "INVALID_REQUEST"
	InternalError  Code = "INTERNAL_ERROR"
)

type Error struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func New(code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}