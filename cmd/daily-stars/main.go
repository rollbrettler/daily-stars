package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	e "github.com/rollbrettler/daily-stars/errors"
	"github.com/rollbrettler/daily-stars/stars"
	"github.com/rollbrettler/daily-stars/username"
)

var port string
var apiUser string
var token string
var rateLimit int
var lastAction time.Time

func init() {
	flag.StringVar(&port, "port", ":8001", "Port to listen on")
	flag.StringVar(&apiUser, "apiUser", "", "GitHub Username for an authorized request")
	flag.StringVar(&token, "token", "", "GitHub token for an authorized request")
}

func main() {
	parseConfigFlags()

	http.HandleFunc("/", showStar)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.ListenAndServe(port, nil)
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func showStar(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ping" {
		status(w)
		return
	}

	username, suffix := username.WithSuffix(r.URL.Path)
	if username == "" {
		jsonErrorResonse(w, e.NoUsername)
		return
	}
	log.Printf("Username: %v\n", username)

	s := stars.New(username, apiUser, token)

	repos, err := s.Repos()
	if err != (e.ResponseError{}) {
		jsonErrorResonse(w, err)
		return
	}

	rateLimit = s.RateLimit
	lastAction = time.Now().UTC()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch suffix {
	case "json":
		jsonResponse(w, repos)
	default:
		jsonResponse(w, repos)
	}
}

func jsonResponse(w http.ResponseWriter, r []stars.StaredRepos) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	marshaledJSON, err := json.Marshal(r)
	if err != nil {
		jsonErrorResonse(w, e.WrongUsername)
		return
	}
	w.Write(marshaledJSON)
}

func jsonErrorResonse(w http.ResponseWriter, err e.ResponseError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	errorJSON, _ := json.Marshal(err)
	w.Write(errorJSON)
}

func parseConfigFlags() {
	flag.Parse()

	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}

	if envAPIUser := os.Getenv("APIUSER"); envAPIUser != "" {
		apiUser = envAPIUser
	}

	if envToken := os.Getenv("TOKEN"); envToken != "" {
		token = envToken
	}
}

func status(w http.ResponseWriter) {
	type status struct {
		RateLimit  int       `json:"rate_limit"`
		LastAction time.Time `json:"last_action"`
	}

	marshaledJSON, err := json.Marshal(status{
		RateLimit:  rateLimit,
		LastAction: lastAction,
	})
	if err != nil {
		jsonErrorResonse(w, e.WrongUsername)
		return
	}
	w.Write(marshaledJSON)
	return
}
