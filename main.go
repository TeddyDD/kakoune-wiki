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

	flag.StringVar(&flagCompleteMediawikiLink, "complete-mediawiki", "", "prefix")
	flag.StringVar(&flagCompleteMarkdownLink, "complete-md-link", "", "prefix")
	flag.Parse()

	err := os.Chdir(cfg.WikiPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch {
	case flagCompleteMediawikiLink != "":
		completeMediawikiCmd(cfg, flagCompleteMediawikiLink)
	case flagCompleteMarkdownLink != "":
		completeMarkdownLinkCmd(cfg, flagCompleteMarkdownLink)
	}
}
