package stars

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

const (
	pageCount = 1
	apiPath   = "https://api.github.com/users/%v/starred?per_page=" + pageCount
)

// Stars …
type Stars struct {
	Pages    int
	Username string
	stared   []StaredRepos
}

// StaredRepos …
type StaredRepos struct {
	Name        string `json:"name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
}

// Repos …
func (s *Stars) Repos() ([]StaredRepos, error) {

	if err := s.setPagesCount(); err != nil {
		return nil, err
	}

	r, err := s.starsFromPage(1)
	if err != nil {
		return nil, err
	}
	s.stared = append(s.stared, r...)

	return s.stared, nil
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
	var r []StaredRepos

	apiURL := fmt.Sprintf(apiPath+"&page=%v", s.Username, strconv.Itoa(p))
	log.Printf("%v\n", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}

	return r, nil
}

func pagesCount(s string) (int, error) {
	if s == "" {
		return 1, nil
	}
	re, err := regexp.Compile("page=\\d+")
	if err != nil {
		return 0, err
	}
	foundStrings := re.FindAllString(s, -1)
	found := foundStrings[len(foundStrings)-1]

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
