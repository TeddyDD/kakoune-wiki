package app

import (
	"github.com/TeddyDD/kakoune-wiki/domain/markdown"
	"github.com/TeddyDD/kakoune-wiki/domain/mediawiki"
)

func (a app) ConvertMediaWikiLinkToMarkdown(link string) (string, error) {
	mediaWikiLink := mediawiki.New(link)

	buffile, err := a.validBuffile()
	if err != nil {
		return "", err
	}

	markdownLink := markdown.NewFrom(
		mediaWikiLink.Alt(),
		a.wiki.AddresToMarkdown(mediaWikiLink.Addres(), buffile),
	)

	return markdownLink.String(), nil
}

func (a app) ConvertMarkdownLinkToMediawiki(link string) (string, error) {
	markdownLink, err := markdown.New(link)
	if err != nil {
		return "", err
	}

	buffile, err := a.validBuffile()
	if err != nil {
		return "", err
	}

	mediaWikiLink := mediawiki.NewFrom(
		markdownLink.Alt(),
		a.wiki.AddresToMediaWiki(markdownLink.Addres(), buffile),
	)
	return mediaWikiLink.String(), nil
}
