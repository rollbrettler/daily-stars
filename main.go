package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type staredRepos struct {
	Name        string `json:"name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
}

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

	username := strings.SplitN(r.URL.Path, "/", 3)[1]
	fmt.Printf("%v\n", username)

	resp, err := http.Get("https://api.github.com/users/" + username + "/starred")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("%v\n\n", resp.Header.Get("Link"))

	var repos []staredRepos

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		fmt.Printf("%v\n", err)
		w.Write([]byte("Wrong username"))
		return
	}

	t, _ := template.ParseFiles("html/index.html")

	t.Execute(w, repos)
}
