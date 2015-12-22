package errors

import (
	"encoding/json"
	"testing"
)

func TestWrongUsernameJsonMarsheling(t *testing.T) {
	_, err := json.Marshal(WrongUsername)
	if err != nil {
		t.Error("Expected no error from json.Marshal() got: ", err)
	}
}

func TestNoUsernameJsonMarsheling(t *testing.T) {
	_, err := json.Marshal(NoUsername)
	if err != nil {
		t.Error("Expected no error from json.Marshal() got: ", err)
	}
}
