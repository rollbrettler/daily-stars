package main

import (
	"flag"
	"net/http"
	"text/template"

	"github.com/rollbrettler/daily-stars/stars"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8001", "Port to listen on")
}

func main() {
	flag.Parse()

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
	var s stars.Stars
	s.URL = r.URL

	repos, err := s.Repos()
	if err != nil {
		w.Write([]byte("Wrong username"))
	}

	t, _ := template.ParseFiles("html/index.html")

	t.Execute(w, repos)
}
