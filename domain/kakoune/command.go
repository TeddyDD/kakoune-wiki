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
