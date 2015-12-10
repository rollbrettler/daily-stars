package errors

import (
	"encoding/json"
	"testing"
)

func TestWrongUsernameJsonMarsheling(t *testing.T) {
	_, err := json.Marshal(WrongUsername)
	if err != nil {
		t.Error("Expected no error from json.Marshal() got ", err)
	}
}
