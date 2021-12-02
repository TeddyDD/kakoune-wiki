package main

import (
	"flag"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
)

var (
	flagInit     bool
	flagRelative bool

	flagConvertToMd        bool
	flagConvertToMediaWiki bool

	flagCompleteMediawikiLink string
	flagCompleteMarkdownLink  string
)

func requireMinArgs(min int, msg string) {
	if len(flag.Args()) < min {
		log.Fatal(msg)
	}
}

func main() {
	cfg := &config{}
	env.Parse(cfg)

	flag.BoolVar(&flagInit, "init", false, "print Kakoune init script")
	flag.BoolVar(&flagRelative, "relative-path", false, "show relative path between DST and SRC")
	flag.BoolVar(&flagConvertToMd, "convert-to-md", false, "convert mediawiki link from stdin to md link")
	flag.BoolVar(&flagConvertToMediaWiki, "convert-to-mediawiki", false, "convert md linkt form stdin to mediawiki")

	flag.StringVar(&flagCompleteMediawikiLink, "complete-mediawiki", "", "prefix")
	flag.StringVar(&flagCompleteMarkdownLink, "complete-md-link", "", "prefix")
	flag.Parse()

	err := os.Chdir(cfg.WikiPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	cfg.cwd = cwd

	switch {
	case flagCompleteMediawikiLink != "":
		completeMediawikiCmd(cfg, flagCompleteMediawikiLink)
	case flagCompleteMarkdownLink != "":
		completeMarkdownLinkCmd(cfg, flagCompleteMarkdownLink)
	case flagConvertToMd:
		convertToMdCmd(cfg)
	case flagConvertToMediaWiki:
		err := convertToMediawikiCmd(cfg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
