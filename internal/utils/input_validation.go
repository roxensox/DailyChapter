package utils

import (
	"slices"
	"strings"
	"time"
)

func ValidateEmail(email string) bool {
	// Takes in an email string and checks if it's a valid address (needs improvement)

	// Returns false immediately if the email was left blank
	if email == "" {
		return false
	}

	// Slice of valid domain extensions
	validDomains := []string{
		"com",
		"gov",
		"org",
		"edu",
		"io",
	}

	// Splits the email string by the @
	sections := strings.Split(email, "@")
	if len(sections) != 2 {
		return false
	}

	// Spits the latter section by the "." to get the domain name
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
