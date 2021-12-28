package app

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/TeddyDD/kakoune-wiki/domain/common"
	"github.com/TeddyDD/kakoune-wiki/domain/kakoune"
	"github.com/TeddyDD/kakoune-wiki/domain/wiki"
)

// TODO: fuzzy match

type CompleteFunc func(string) (kakoune.Completions, error)

func (a app) RunCompleter(fn CompleteFunc, prefix string) {
	completions, err := fn(prefix)
	if err != nil {
		return
	}
	cmd := kakoune.SetCompletions(a.config.Client, completions)
	fmt.Print(cmd)
}

func (a app) CompleteMediaWiki(prefix string) (kakoune.Completions, error) {
	res := make(kakoune.Completions, 0)

	files, err := a.files()
	if err != nil {
		return nil, fmt.Errorf("coundn't get list of files in wiki: %w", err)
	}
	files = wiki.FilterMarkdown(files)

	for i := range files {
		relative, err := a.wiki.RelativeToWiki(files[i])
		if err != nil {
			return nil, err
		}

		if strings.Contains(strings.ToLower(relative), strings.ToLower(prefix)) {
			res = append(res, kakoune.NewCompletionEntry(
				common.TrimExtension(relative, filepath.Ext(relative)),
				relative,
			))
		}
	}

	return res, nil
}

func (a app) CompleteMarkdown(prefix string) (kakoune.Completions, error) {
	completions := kakoune.Completions{}
	_, err := a.buffileInWiki()
	if err != nil {
		return nil, err
	}

	files, err := a.files()
	if err != nil {
		return nil, fmt.Errorf("coundn't get list of files in wiki: %w", err)
	}
	files = wiki.FilterMarkdown(files)

	for i := range files {
		rel, err := common.RelativePath(files[i], a.config.Buffile)
		if err != nil {
			return nil, err
		}

		relToWikiRoot, err := a.wiki.RelativeToWiki(files[i])
		if err != nil {
			return nil, err
		}

		if strings.Contains(
			strings.ToLower(relToWikiRoot),
			strings.ToLower(prefix),
		) {
			completions = append(
				completions,
				kakoune.NewCompletionEntry(
					rel,
					relToWikiRoot,
				),
			)
		}

	}
	return completions, nil
}
