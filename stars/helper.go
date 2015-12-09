package stars

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func pagesCount(s string) (int, error) {
	if s == "" {
		return 1, nil
	}

	m := make(map[string]int)

	pages := regexp.MustCompile(`.*[\?&]page\=(?P<page>\d+).*rel=\"(?P<rel>.*)\"`)

	for _, link := range strings.Split(s, ",") {
		var rel string
		var number int

		match := pages.FindStringSubmatch(link)

		for i, name := range pages.SubexpNames() {
			switch name {
			case "rel":
				rel = match[i]
			case "page":
				number, _ = strconv.Atoi(match[i])
			}
		}

		if rel != "" && number > 0 {
			m[rel] = number
		}
	}

	if m["last"] == 0 && m["prev"] > 0 {
		return m["prev"] + 1, nil
	}
	return m["last"], nil
}

func randomPageNumber(i int) int {
	const shortForm = "2006-January-02"

	year, month, day := time.Now().Date()

	date := fmt.Sprintf("%v-%v-%v", year, month, day)
	t, _ := time.Parse(shortForm, date)
	rand.Seed(t.Unix())

	randomNumber := rand.Intn(i)
	log.Printf("Random: %v\n", randomNumber)

	return randomNumber
}
