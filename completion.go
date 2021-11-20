package main

import (
	"fmt"
	"strings"
)

type completion struct {
	item    string
	command string
	hint    string
}

func (c completion) String() string {
	return fmt.Sprintf("%s|%s|%s", c.item, c.command, c.hint)
}

func printCompletion(cfg *config, completions []string) {
	fmt.Printf(
		"evaluate-commands -verbatim -try-client '%s' set-option -add buffer wiki_completions %s",
		cfg.Client,
		strings.Join(completions, " "),
	)
}
