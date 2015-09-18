package stars

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	apiEnpoint = "https://api.github.com/users/"
	apiPath    = "/starred"
)

// Stars …
type Stars struct {
	URL      *url.URL
	Pages    int64
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

	r, err := s.starsFromPage(1)
	if err != nil {
		return nil, err
	}
	repos = append(repos, r...)

	return repos, nil
}

func (s *Stars) starsFromPage(p int64) ([]StaredRepos, error) {
	var repos []StaredRepos

	apiURL := apiEnpoint + s.Username + apiPath + "?page=" + strconv.Itoa(int(p))

	log.Printf("%v\n", apiURL)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	s.Pages, err = pagesCount(resp.Header.Get("Link"))
	log.Printf("%v\n", s.Pages)

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}

	return repos, nil
}

func pagesCount(s string) (int64, error) {
	re, err := regexp.Compile("page=\\d+")
	if err != nil {
		return 0, err
	}
	found := re.FindAllString(s, -1)

	re2, err := regexp.Compile("\\d+")
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(re2.FindString(found[1]), 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}
