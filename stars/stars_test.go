package stars

import (
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	var s Stars
	s = New("rollbrettler")

	if s.Username != "rollbrettler" {
		t.Error("Expected 'rollbrettler', got ", s.Username)
	}
	if s.Pages != 0 {
		t.Error("Expected 0 Pages, got ", s.Pages)
	}
}

func TestRepos(t *testing.T) {
	var s Stars
	s = New("octocat")
	r, err := s.Repos()

	if err != nil {
		t.Error("Expected no errors while fetching the repos, got ", err)
		return
	}

	count, _ := strconv.Atoi(pageCount)
	if len(r) != count {
		t.Error("Expected results, got ", len(r))
		return
	}

	if s.Pages != 2 {
		t.Error("Expected to fetch pages, got", s.Pages)
		return
	}
}
