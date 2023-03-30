package helpers

import "regexp"

func IsValidURL(url string) bool {
	pattern := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`

	regexp := regexp.MustCompile(pattern)

	return regexp.MatchString(url)
}
