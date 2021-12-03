package markdown

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Link struct {
	alt    string
	addres string
}

func (l Link) Alt() string {
	return l.alt
}

func (l Link) Addres() string {
	return l.addres
}

var (
	protocolRegexp = regexp.MustCompile(`^[\w]+:[/]{2}`)
	linkRegexp     = regexp.MustCompile(`\[(.*)\]\((.*)\)`)

	WrongInput = errors.New("wrong input")
)

// IsInternal checks if this is internal wiki link
func (l Link) IsInternal() bool {
	// no absolute path
	if strings.HasPrefix(l.addres, "/") {
		return false
	}
	// no protocol
	if protocolRegexp.MatchString(l.addres) {
		return false
	}

	return true
}

func (l Link) String() string {
	return fmt.Sprintf("[%s](%s)", l.alt, l.addres)
}

func New(in string) (link Link, err error) {
	groups := linkRegexp.FindStringSubmatch(in)
	if groups == nil || len(groups) != 3 {
		return link, fmt.Errorf("input: '%s': %s", in, WrongInput)
	}

	link.alt = groups[1]
	link.addres = groups[2]

	if link.alt == "" {
		link.alt = link.addres
	}
	return link, nil
}

func NewFrom(alt string, addr string) Link {
	return Link{
		alt:    alt,
		addres: addr,
	}
}
