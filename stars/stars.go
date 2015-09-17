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

type Stars struct {
}

type staredRepos struct {
	Name        string `json:"name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
}

func (s *Stars) Repos(url *url.URL) (repos []staredRepos, err error) {
	username := strings.SplitN(url.Path, "/", 3)[1]
	log.Printf("%v\n", username)

	resp, err := http.Get("https://api.github.com/users/" + username + "/starred")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pages, err := getPagesCount(resp.Header.Get("Link"))
	log.Printf("%v\n", pages)

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}

	return repos, nil
}

func getPagesCount(s string) (i int64, err error) {
	re, err := regexp.Compile("page=\\d+")
	if err != nil {
		return 0, err
	}
	found := re.FindAllString(s, -1)

	re2, err := regexp.Compile("\\d+")
	if err != nil {
		return 0, err
	}

	i, err = strconv.ParseInt(re2.FindString(found[1]), 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}
