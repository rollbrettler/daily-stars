package stars

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	e "github.com/rollbrettler/daily-stars/errors"
)

const (
	pageCount = "1"
	apiPath   = "https://api.github.com/users/%v/starred?per_page=" + pageCount
)

// Stars is the returned struct
type Stars struct {
	Pages    int
	Username string
	stared   []StaredRepos
}

// StaredRepos is a struct to unmarshal the json response
type StaredRepos struct {
	Name        string `json:"name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
}

// New returns a new stars struct with the given username
func New(username string) Stars {
	return Stars{
		Username: username,
	}
}

// Repos returns a slice of StaredRepos
func (s *Stars) Repos() ([]StaredRepos, e.ResponseError) {

	if err := s.setPagesCount(); err != nil {
		log.Printf("Error in Repos after calling setPagesCount: %v\n", err)
		return nil, e.Unhandled
	}

	c1 := make(chan []StaredRepos, 1)
	errCh := make(chan e.ResponseError, 1)
	go func() {
		r, err := s.starsFromPage(randomPageNumber(s.Pages))
		if err != nil {
			errCh <- e.NoUsername
			return
		}
		c1 <- r
	}()

	select {
	case res := <-c1:
		s.stared = append(s.stared, res...)
		return s.stared, e.ResponseError{}
	case <-time.After(time.Second * 5):
		return nil, e.TimeOut
	case err := <-errCh:
		return nil, err
	}
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
	log.Printf("%v Pages\n", s.Pages)

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
