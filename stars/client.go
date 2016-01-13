package stars

import (
	"log"
	"net/http"
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

	log.Printf("Remaining requests: %v\n", resp.Header.Get("X-Ratelimit-Remaining"))

	return resp, nil
}
