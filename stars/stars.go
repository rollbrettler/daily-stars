package stars

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	apiPath = "https://api.github.com/users/%v/starred"
)

// Stars …
type Stars struct {
	URL      *url.URL
	Pages    int
	Username string
}

// StaredRepos …
type StaredRepos struct {
	Name        string `json:"name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
}

// Repos …
func (s *Stars) Repos() ([]StaredRepos, error) {
	var repos []StaredRepos

	s.Username = strings.SplitN(s.URL.Path, "/", 3)[1]
	log.Printf("%v\n", s.Username)

	if err := s.setPagesCount(); err != nil {
		return nil, err
	}

	r, err := s.starsFromPage(1)
	if err != nil {
		return nil, err
	}
	repos = append(repos, r...)

	return repos, nil
}

func (s *Stars) setPagesCount() error {
	apiURL := fmt.Sprintf(apiPath, s.Username)
	log.Printf("%v\n", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	s.Pages, err = pagesCount(resp.Header.Get("Link"))
	log.Printf("%v\n", s.Pages)

	return nil
}

func (s *Stars) starsFromPage(p int) ([]StaredRepos, error) {
	var repos []StaredRepos

	apiURL := fmt.Sprintf(apiPath+"?page=%v", s.Username, strconv.Itoa(p))
	log.Printf("%v\n", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}

	return repos, nil
}

func pagesCount(s string) (int, error) {
	if s == "" {
		return 1, nil
	}
	re, err := regexp.Compile("page=\\d+")
	if err != nil {
		return 0, err
	}
	found := re.FindAllString(s, -1)[1]

	re2, err := regexp.Compile("\\d+")
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(re2.FindString(found))
	if err != nil {
		return 0, err
	}

	return i, nil
}
