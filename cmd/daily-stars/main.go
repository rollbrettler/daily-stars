package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	e "github.com/rollbrettler/daily-stars/errors"
	"github.com/rollbrettler/daily-stars/stars"
	"github.com/rollbrettler/daily-stars/username"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8001", "Port to listen on")
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
	username, suffix := username.WithSuffix(r.URL.Path)
	if username == "" {
		jsonErrorResonse(w, e.NoUsername)
		return
	}
	log.Printf("Username: %v\n", username)

	s := stars.New(username)

	repos, err := s.Repos()
	if err != (e.ResponseError{}) {
		jsonErrorResonse(w, err)
		return
	}

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

	envPort := os.Getenv("PORT")

	if envPort != "" {
		port = ":" + envPort
	}
}
