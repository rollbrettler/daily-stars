package username

import (
	"net/url"
	"testing"
)

var u url.URL
var usernameTest, suffixTest string

func TestUsernameWithoutPrefix(t *testing.T) {
	u, _ := url.Parse("http://example.com/username")
	usernameTest, suffixTest = WithSuffix(u.Path)

	if usernameTest != "username" {
		t.Error("Expected username to be 'username' ", usernameTest)
	}

	if suffixTest != "" {
		t.Error("Expected suffix to be empty ", suffixTest)
	}
}

func TestUsernameWithPrefix(t *testing.T) {
	u, _ := url.Parse("http://example.com/username.prefix")
	usernameTest, suffixTest = WithSuffix(u.Path)

	if usernameTest != "username" {
		t.Error("Expected username to be 'username' ", usernameTest)
	}

	if suffixTest != "prefix" {
		t.Error("Expected suffix to be 'prefix' ", suffixTest)
	}
}
