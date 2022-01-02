package app

import (
	"fmt"
	"path/filepath"

	"github.com/TeddyDD/kakoune-wiki/domain/kakoune"
)

func (a app) EditMarkdown(link string) string {
	_, err := a.buffileInWiki()
	if err != nil {
		return kakoune.Fail(err.Error())
	}

	path := filepath.Join(a.config.Buffile, "..", link)
	return kakoune.Edit(a.config.Client, path)
}
