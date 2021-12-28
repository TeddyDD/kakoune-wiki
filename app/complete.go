package app

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/TeddyDD/kakoune-wiki/domain/kakoune"
	"github.com/TeddyDD/kakoune-wiki/domain/wiki"
)

// TODO: fuzzy match
// TODO: match path not only base name?

type CompleteFunc func(string) (kakoune.Completions, error)

func (a app) CompleteMediaWiki(prefix string) (kakoune.Completions, error) {
	res := make(kakoune.Completions, 0)

	files, err := a.files()
	if err != nil {
		return nil, fmt.Errorf("coundn't get list of files in wiki: %w", err)
	}

	files = wiki.FilterMarkdown(files)
	files = wiki.FilterPrefixNoCase(files, prefix)

	for i := range files {
		relative, err := a.wiki.RelativeToWiki(files[i])
		if err != nil {
			return nil, err
		}

		res = append(res, kakoune.NewCompletionEntry(
			strings.TrimSuffix(relative, filepath.Ext(relative)),
			relative,
		))
	}

	return res, nil
}
