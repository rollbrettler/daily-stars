package username

import "regexp"

// WithSuffix gets the username from an url path of type string
// including the suffix
func WithSuffix(urlPath string) (string, string) {
	var username, suffix string
	regex := regexp.MustCompile(`\/(?P<username>[^.\s]*)\.?(?P<suffix>[^.\s]*)`)
	match := regex.FindStringSubmatch(urlPath)

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
