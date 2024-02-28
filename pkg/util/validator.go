package util

import (
	"regexp"
	"time"
)

const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,6}$`

// usernameRegex is a regex to validate username with the following rules:
// - Minimum 3 characters
// - Only alphanumeric, underscore
const usernameRegex = `^[a-zA-Z0-9_]+$`
const locationRegex = `^[A-Za-z' -]+$`
const dateFormat = "2006-01-02" // ISO 8601 date format

func HasSpace(s string) bool {
	for _, c := range s {
		if c == ' ' {
			return true
		}
	}
	return false
}

func IsValidUsername(s string) bool {
	re := regexp.MustCompile(usernameRegex)
	return !HasSpace(s) && !re.MatchString(s)
}

func IsValidLocation(s string) bool {
	if len(s) <= 3 {
		return false
	}

	re := regexp.MustCompile(locationRegex)
	return re.MatchString(s)
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func IsValidMinLength(s string, min int) bool {
	return len(s) >= min
}

func IsValidDateOfBirth(s string) bool {
	_, err := time.Parse(dateFormat, s)
	return err == nil
}
