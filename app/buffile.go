package app

import "errors"

// buffileInWiki returns path to buffile relative to wiki root
func (a app) buffileInWiki() (string, error) {
	if a.config.Buffile == "" {
		return "", errors.New("buffile not set")
	}

	if !a.wiki.FileInWiki(a.config.Buffile) {
		return "", errors.New("buffile outisde of wiki path")
	}

	return a.wiki.RelativeToWiki(a.config.Buffile)
}
