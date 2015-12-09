package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"text/template"

	"github.com/rollbrettler/daily-stars/stars"
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

	username, suffix := username(r.URL)
	if username == "" {
		t, _ := template.ParseFiles("html/index.html")
		t.Execute(w, r.Host)
		return
	}
	log.Printf("Username: %v\n", username)

	s := stars.New(username)

	repos, err := s.Repos()
	if err != nil {
		w.Write([]byte("Wrong username"))
		return
	}

	t, _ := template.ParseFiles("html/result.html")

	if suffix == "json" {
		jsonResponse(w, repos)
	} else {
		t.Execute(w, repos)
	}

}

func jsonResponse(w http.ResponseWriter, r []stars.StaredRepos) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	marshaledJSON, err := json.Marshal(r)
	if err != nil {
		w.Write([]byte("{'error': 'Wrong username'}"))
	}
	w.Write(marshaledJSON)
}

func username(url *url.URL) (string, string) {
	var username, suffix string
	regex := regexp.MustCompile(`\/(?P<username>[^.\s]*)\.?(?P<suffix>[^.\s]*)`)
	match := regex.FindStringSubmatch(url.Path)

	for i, name := range regex.SubexpNames() {
		switch name {
		case "username":
			username = match[i]
		case "suffix":
			suffix = match[i]
		}
	}

	return username, suffix
}

func parseConfigFlags() {
	flag.Parse()

	envPort := os.Getenv("PORT")

	if envPort != "" {
		port = ":" + envPort
	}
}
