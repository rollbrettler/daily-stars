package errors

// ResponseError to be rendered as a error response
type ResponseError struct {
	Message string
}

func (e ResponseError) Error() string {
	return e.Message
}

// Unhandled response if the error is not yet handled correctly
var Unhandled = ResponseError{Message: "Unhandled error"}

// Unhandled response if the error is not yet handled correctly
var TimeOut = ResponseError{Message: "Time out while connecting to remote host"}

// WrongUsername response if the username cannot be recognized
var WrongUsername = ResponseError{Message: "Wrong username"}

// NoUsername response if there is no username given
var NoUsername = ResponseError{Message: "No username. Please use this format https://%url%/%username%"}
