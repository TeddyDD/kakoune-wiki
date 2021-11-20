package main

import "strings"

func normalizedContains(str, sub string) bool {
	str = strings.ToLower(str)
	sub = strings.ToLower(sub)
	return strings.Contains(str, sub)
}
