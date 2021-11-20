package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func cleanPrefix(prefix string) string {
	// completion is case insensitive
	prefix = strings.ToLower(prefix)

	prefix = strings.TrimPrefix(prefix, "[[")
	prefix = strings.TrimSuffix(prefix, "]]")
	if idx := strings.LastIndex(prefix, "|"); idx > 0 {
		prefix = prefix[idx+1:]
	}

	return prefix
}

func completeMediawikiCmd(cfg *config, prefix string) error {
	completions := []string{}
	prefix = cleanPrefix(prefix)

	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		lowerPath := strings.ToLower(path)

		if strings.Contains(lowerPath, prefix) &&
			strings.ToLower(filepath.Ext(lowerPath)) == ".md" {

			completions = append(completions, `%{`+completion{
				item: strings.TrimSuffix(path, filepath.Ext(path)),
				hint: path,
			}.String()+`}`)
		}
		return nil
	})

	printCompletion(cfg, completions)

	return nil
}
