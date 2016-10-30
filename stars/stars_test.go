package stars

import (
	"strconv"
	"testing"

	e "github.com/rollbrettler/daily-stars/errors"
)

func TestNew(t *testing.T) {
	var s Stars
	s = New("rollbrettler", "", "")

	if s.Username != "rollbrettler" {
		t.Error("Expected 'rollbrettler', got: ", s.Username)
	}
	if s.Pages != 0 {
		t.Error("Expected 0 Pages, got: ", s.Pages)
	}
}

func TestRepos(t *testing.T) {
	var s Stars
	s = New("octocat", "", "")
	r, err := s.Repos()

	if err != (e.ResponseError{}) {
		t.Error("Expected no errors while fetching the repos, got: ", err)
		return
	}

	count, _ := strconv.Atoi(pageCount)
	if len(r) != count {
		t.Error("Expected results, got ", len(r))
		return
	}

	if s.Pages != 3 {
		t.Error("Expected to fetch 3 pages, got: ", s.Pages)
		return
	}
}

func TestReposWrongUsername(t *testing.T) {
	var s Stars
	s = New("UsernameThatDoesNotExist", "", "")
	r, err := s.Repos()

	if err != e.WrongUsername {
		t.Error("Expected to return error for wrong username got: ", err)
		return
	}

	if r != nil {
		t.Error("Expected to return empty []StaredRepos got: ", r)
		return
	}
}
