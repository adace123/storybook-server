package utils

import (
	"regexp"
)

func ValidEmail(email string) bool {
	return regexp.MustCompile("^((\\w+[\\.]?)+)@(\\w+\\.){1,}\\w{2,9}$").MatchString(email)
}
