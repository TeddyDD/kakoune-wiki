package main

import (
	"flag"
)

var (
	cmdConvertLink  = flag.Bool("convert-link", false, "link conventer")
	flagToMarkdown  = flag.Bool("to-markdown", false, "convert link to markdown format")
	flagToMediawiki = flag.Bool("to-mediawiki", false, "convert link to mediawiki format")

	cmdComplete               = flag.Bool("complete", false, "autocompletion")
	flagCompleteMarkdownLink  = flag.String("markdown-prefix", "", "prefix for completion")
	flagCompleteMediawikiLink = flag.String("mediawiki-prefix", "", "prefix for completion")
)

func setupFlags() {
	flag.Parse()
}
