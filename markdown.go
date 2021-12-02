package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type mdLink struct {
	alt    string
	addres string
}

var protocolRegexp = regexp.MustCompile(`^[\w]+:[/]{2}`)

func (l mdLink) isInternal() bool {
	// no absolute path
	if strings.HasPrefix(l.addres, "/") {
		return false
	}
	// no protocol
	if protocolRegexp.MatchString(l.addres) {
		return false
	}

	return true
}

func (l mdLink) String() string {
	return fmt.Sprintf("[%s](%s)", l.alt, l.addres)
}

var linkRegexp = regexp.MustCompile(`\[(.*)\]\((.*)\)`)

func newMdLink(in string) (link mdLink, err error) {
	groups := linkRegexp.FindStringSubmatch(in)
	if groups == nil || len(groups) != 3 {
		return link, errors.New("wrong input")
	}

	link.alt = groups[1]
	link.addres = groups[2]

	if link.alt == "" {
		link.alt = link.addres
	}
	return link, nil
}

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

func convertToMdCmd(cfg *config) error {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)
	var line string
	if scan.Scan() {
		line = scan.Text()
	} else {
		return errors.New("no input")
	}

	link := newMediaWikiLink(line)

	targetMd := fmt.Sprintf("%s.md", link.addres)
	targetMd, err := filepath.Abs(targetMd)
	if err != nil {
		return nil
	}
	thisMd := cfg.Buffile

	res, err := relativePath(targetMd, thisMd)
	if err != nil {
		return nil
	}

	fmt.Printf("[%s](%s)", link.alt, res)
	return nil
}
