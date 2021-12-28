package app

import (
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
