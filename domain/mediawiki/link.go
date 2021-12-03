package mediawiki

import (
	"fmt"
	"strings"
)

type Link struct {
	addres string
	alt    string
}

// Addres returns addres of given link
func (l Link) Addres() string {
	return l.addres
}

// Alt returns alternative text for given link
func (l Link) Alt() string {
	return l.alt
}

func (l Link) String() string {
	if l.alt != "" {
		return fmt.Sprintf("[[%s|%s]]", l.alt, l.addres)
	}
	return fmt.Sprintf("[[%s]]", l.addres)
}

func New(from string) Link {
	from = strings.TrimSpace(from)
	from = strings.TrimPrefix(from, "[[")
	from = strings.TrimSuffix(from, "]]")
	parts := strings.Split(from, "|")
	l := len(parts)
	switch {
	case l == 0:
		return Link{}
	case l == 1:
		return Link{
			addres: parts[0],
			alt:    parts[0],
		}
	default:
		return Link{
			addres: strings.Join(parts[1:], ""),
			alt:    parts[0],
		}
	}
}

func NewFrom(alt, addres string) Link {
    return Link{
    	addres: addres,
    	alt:    alt,
    }
}
