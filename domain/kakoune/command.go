package kakoune

import (
	"fmt"
)

func SetCompletions(client string, completion Completions) string {
	return fmt.Sprintf(
		"evaluate-commands -verbatim -try-client '%s' set-option -add buffer wiki_completions %s",
		client,
		completion,
	)
}

func Edit(client, path string) string {
	return fmt.Sprintf(
		"evaluate-commands -client %s %%{ edit '%s'}",
		client,
		EscapeSingleQuote(path),
	)
}

func Fail(msg string) string {
	return fmt.Sprintf(
		"fail '%s'",
		EscapeSingleQuote(msg),
	)
}

func Debug(msg string) string {
	return fmt.Sprintf(
		"echo -debug wiki: '%s'",
		EscapeSingleQuote(msg),
	)
}
