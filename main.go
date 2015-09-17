package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	log.Printf("%v\n", username)

	resp, err := http.Get("https://api.github.com/users/" + username + "/starred")
	if err != nil {
		log.Print(err)
		w.Write([]byte("Server error\n"))
		return
	}
	defer resp.Body.Close()

	pages, err := getPagesCount(resp.Header.Get("Link"))
	log.Printf("%v\n", pages)

	var repos []staredRepos

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Printf("%v\n", err)
		w.Write([]byte("Wrong username"))
		return
	}

	t, _ := template.ParseFiles("html/index.html")

	t.Execute(w, repos)
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
