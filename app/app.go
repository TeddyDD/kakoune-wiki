package app

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/TeddyDD/kakoune-wiki/domain/kakoune"
	"github.com/TeddyDD/kakoune-wiki/domain/wiki"
)

type app struct {
	config *kakoune.Config
	wiki   *wiki.Wiki

	// files returns list of files in wiki dir
	files func() ([]string, error)
}

func New(
	config *kakoune.Config,
	wiki *wiki.Wiki,
) app {
	if config == nil {
		panic("no config")
	}
	if wiki == nil {
		panic("no wiki")
	}

	return app{
		config: config,
		wiki:   wiki,

		files: func() ([]string, error) {
			return wiki.Files()
		},
	}
}

var singleQuote = regexp.MustCompile(`'`)

func (a app) Debug(msg string) {
	if a.config == nil || !a.config.Debug {
		return
	}
	kak, err := exec.LookPath("kak")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(kak, "-p", a.config.Session)

	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer in.Close()

	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	msg = singleQuote.ReplaceAllString(msg, `''`)

	fmt.Fprintf(in, "echo -debug wiki: '%s'", msg)
}

func (a app) Debugf(format string, val ...interface{}) {
	a.Debug(fmt.Sprintf(format, val...))
}
