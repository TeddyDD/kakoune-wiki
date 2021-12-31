package main

import (
	"flag"
)

var (
	cmdConvertLink  = flag.Bool("convert-link", false, "link conventer")
	flagToMarkdown  = flag.Bool("to-markdown", false, "convert link to markdown format")
	flagToMediawiki = flag.Bool("to-mediawiki", false, "convert link to mediawiki format")

	cmdComplete               = flag.Bool("complete", false, "autocompletion")
	flagCompleteWikiCmd       = flag.Bool("wiki-cmd", false, "args passed to wiki cmd")
	flagAllFiles              = flag.Bool("all-markdown-files", false, "list all Markdown files in wiki")
	flagCompleteMarkdownLink  = flag.String("markdown-prefix", "", "prefix for completion")
	flagCompleteMediawikiLink = flag.String("mediawiki-prefix", "", "prefix for completion")
)

func setupFlags() {
	flag.Parse()
}
