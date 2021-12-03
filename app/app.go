package app

import (
	"github.com/TeddyDD/kakoune-wiki/domain/kakoune"
	"github.com/TeddyDD/kakoune-wiki/domain/wiki"
)

type app struct {
	config *kakoune.Config
	wiki   *wiki.Wiki
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
	}
}
