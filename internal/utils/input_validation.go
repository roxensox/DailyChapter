package utils

import (
	"slices"
	"strings"
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
