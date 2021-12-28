package common

import "strings"

// Contains checks if str contains substring (case insensitive)
func Contains(str, sub string) bool {
	str = strings.ToLower(str)
	sub = strings.ToLower(sub)
	return strings.Contains(str, sub)
}
