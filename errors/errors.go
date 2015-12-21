package errors

// ResponseError to be rendered as a error response
type ResponseError struct {
	Message string
}

// WrongUsername response if the username cannot be recognized
var WrongUsername = ResponseError{Message: "Wrong username"}