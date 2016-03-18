package stars

import (
	"log"
	"net/http"
	"strconv"
)

func (s *Stars) apiGetRequest(apiURL string) (*http.Response, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", apiURL, nil)
	if s.apiUser != "" && s.token != "" {
		req.SetBasicAuth(s.apiUser, s.token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	s.RateLimit, err = strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		s.RateLimit = 0
	}
	log.Printf("Remaining requests: %v\n", s.RateLimit)

	return resp, nil
}
