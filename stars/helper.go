package stars

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func pagesCount(s string) (int, error) {
	if s == "" {
		return 1, nil
	}
	re, err := regexp.Compile("page=\\d+")
	if err != nil {
		return 0, err
	}
	foundStrings := re.FindAllString(s, -1)
	found := foundStrings[len(foundStrings)-1]

	re2, err := regexp.Compile("\\d+")
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(re2.FindString(found))
	if err != nil {
		return 0, err
	}

	return i, nil
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
