package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func completeMarkdownLinkCmd(cfg *config, link string) error {
	completions := []string{}
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() || strings.ToLower(filepath.Ext(path)) != ".md" {
			return nil
		}

		abs, err := filepath.Abs(path)
		if err != nil {
			return nil
		}
		rel, err := relativePath(abs, cfg.Buffile)
		if err != nil {
			return nil
		}

		if normalizedContains(rel, link) || normalizedContains(path, link) {
			completions = append(completions, `%{`+completion{
				item: rel,
				hint: path,
			}.String()+`}`)
		}

		return nil
	})

	printCompletion(cfg, completions)

	return nil
}
