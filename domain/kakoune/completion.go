package kakoune

import (
	"fmt"
	"strings"
)

type CompletionEntry struct {
	item    string
	command string
	hint    string
}

func NewCompletionEntry(item, hint string) CompletionEntry {
	return CompletionEntry{
		item:    item,
		command: "",
		hint:    hint,
	}
}

func (c CompletionEntry) SetCommand(cmd string) CompletionEntry {
	c.command = cmd
	return c
}

func (c CompletionEntry) String() string {
	return fmt.Sprintf("'%s|%s|%s'", c.item, c.command, c.hint)
}

type Completions []CompletionEntry

func (c Completions) String() string {
	tmp := []string{}

	for i := range c {
		tmp = append(tmp, c[i].String())
	}
	return strings.Join(tmp, " ")
}
