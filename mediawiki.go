package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type mediaWikiLink struct {
	addres string
	alt    string
}

func (l mediaWikiLink) String() string {
	if l.alt != "" {
		return fmt.Sprintf("[[%s|%s]]", l.alt, l.addres)
	}
	return fmt.Sprintf("[[%s]]", l.addres)
}

func newMediaWikiLink(from string) mediaWikiLink {
	from = strings.TrimSpace(from)
	from = strings.TrimPrefix(from, "[[")
	from = strings.TrimSuffix(from, "]]")
	parts := strings.Split(from, "|")
	l := len(parts)
	switch {
	case l == 0:
		return mediaWikiLink{}
	case l == 1:
		return mediaWikiLink{
			addres: parts[0],
			alt:    parts[0],
		}
	default:
		return mediaWikiLink{
			addres: strings.Join(parts[1:], ""),
			alt:    parts[0],
		}
	}
}

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

func convertToMediawikiCmd(cfg *config) error {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)
	var line string
	if scan.Scan() {
		line = scan.Text()
	} else {
		return errors.New("no input")
	}

	link, err := newMdLink(line)
	if err != nil {
		return err
	}

	mwLink := &mediaWikiLink{
		addres: link.addres,
		alt:    link.alt,
	}

	// if internal link then
	if link.isInternal() {
		bufDir := filepath.Dir(cfg.Buffile)
		target := filepath.Join(bufDir, link.addres)
		target = strings.TrimSuffix(target, filepath.Ext(target))

		target, err := filepath.Rel(cfg.cwd, target)
		if err != nil {
			return err
		}

		mwLink.addres = target
		mwLink.alt = strings.TrimSuffix(mwLink.alt, filepath.Ext(mwLink.alt))
	}

	// print converted mediawiki link
	fmt.Print(mwLink)
	return nil
}
