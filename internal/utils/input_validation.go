package utils

import (
	"slices"
	"strings"
	"time"
)

func ValidateEmail(email string) bool {
	validDomains := []string{
		"com",
		"gov",
		"org",
		"edu",
		"io",
	}
	if email == "" {
		return false
	}
	sections := strings.Split(email, "@")
	if len(sections) != 2 {
		return false
	}
	domain := strings.Split(sections[1], ".")
	if !slices.Contains(validDomains, domain[len(domain)-1]) {
		return false
	}
	return true
}

func ParseDate(date string) (time.Time, error) {
	// Accepts a date in the format of YYYY-MM-DD and returns it as a timestamp

	// Establishes parsing pattern
	pattern := "2006-01-02"
	// Parses and returns
	return time.Parse(pattern, date)
}
