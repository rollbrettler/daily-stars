package errors

import (
	"encoding/json"
	"testing"
)

func TestThatResponseErrorActsAsTypeError(t *testing.T) {
	responseError := ResponseError{Message: "error message"}
	if responseError.Error() != "error message" {
		t.Error("Expected error message to be 'error message' got: ", responseError.Error())
	}
}

func TestJsonMarsheling(t *testing.T) {
	if _, err := json.Marshal(ResponseError{Message: "error message"}); err != nil {
		t.Error("Expected no error from json.Marshal() got: ", err)
	}
}
