package utils

import (
	"strings"
)

func IsUrlCorrect(url string) bool {
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		return false
	}

	if strings.Count(url, ".") != 1 {
		return false
	}

	if strings.Contains(url, " ") {
		return false
	}

	return true
}
