package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/TeddyDD/kakoune-wiki/app"
	"github.com/TeddyDD/kakoune-wiki/domain/kakoune"
	"github.com/TeddyDD/kakoune-wiki/domain/wiki"
)

func main() {
	setupFlags()

	config, err := kakoune.FromEnv()
	if err != nil {
		log.Fatal("failed to create config from env")
	}

	err = os.Chdir(config.WikiPath)
	if err != nil {
		log.Fatalf("failed to cd into wiki directory '%s': %s",
			config.WikiPath, err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to read current working directory: %s", err.Error())
	}

	w, err := wiki.New(dir)
	if err != nil {
		log.Fatalf("failed to use %s as wiki directory: %s", dir, err.Error())
	}

	a := app.New(config, w)

	switch {
	case *cmdConvertLink:
		if *flagToMarkdown {
			runFilter(a.ConvertMediaWikiLinkToMarkdown)
		} else {
			runFilter(a.ConvertMarkdownLinkToMediawiki)
		}
	case *cmdComplete:
		if *flagCompleteMarkdownLink != "" {
		} else if *flagCompleteMediawikiLink != "" {
			runCompleter(*flagCompleteMediawikiLink, a.CompleteMediaWiki)
		}
	}
}

func readInput() (string, error) {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)
	var line string
	if scan.Scan() {
		line = scan.Text()
	} else {
		return "", errors.New("no input")
	}
	return line, nil
}

func runCompleter(prefix string, cmd app.CompleteFunc) {
	completions, err := cmd(prefix)
	if err != nil {
		// TODO: report to *debug*
		return
	}

	// TODO: kakoune command or kak -p
	fmt.Println(completions)
}

// run command used as filter in kakoune (for example in |)
// in case of error, prints input or nothing
// TODO: report to *debug*
func runFilter(cmd func(string) (string, error)) {
	in, err := readInput()
	if err != nil {
		return
	}
	out, err := cmd(in)
	if err != nil {
		fmt.Print(in)
		return
	}

	fmt.Print(out)
}
