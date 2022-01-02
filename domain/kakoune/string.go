package kakoune

import "strings"

func EscapeSingleQuote(in string) string {
	return strings.ReplaceAll(in, `'`, `''`)
}
