package main

import (
	"bufio"
    flag "github.com/spf13/pflag"
	"fmt"
	"github.com/muziyoshiz/ansible2tab/formatter"
	"github.com/muziyoshiz/ansible2tab/parser"
	"os"
)

var (
	formatOpt = flag.StringP("format", "f","tsv", `output format "tsv", "js", "md", "md-code", "blg" or "blg-code"`)
	versionOpt = flag.BoolP("version", "v", false, `version`)
)

func main() {
	flag.Parse()

	if *versionOpt {
		fmt.Printf("ansible2tab version: %s (rev: %s)\n", version, revision)
		os.Exit(0)
	}

	var f formatter.Formatter
	switch *formatOpt {
	case "tsv":
		f = &formatter.TSVFormatter{}
	case "js", "json":
		f = &formatter.JSONFormatter{}
	case "md", "markdown":
		f = &formatter.MarkdownFormatter{}
	case "md-code", "markdown-code":
		f = &formatter.MarkdownCodeFormatter{}
	case "blg", "backlog":
		f = &formatter.BacklogFormatter{}
	case "blg-code", "backlog-code":
		f = &formatter.BacklogCodeFormatter{}
	default:
		flag.Usage()
		os.Exit(1)
	}

	parse := parser.Parser()

	fmt.Print(f.GetHeader())

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		res, ok := parse(line)
		if ok {
			fmt.Print(f.Format(res))
		}
	}

	res, ok := parse(parser.EOF)
	if ok {
		fmt.Print(f.Format(res))
	}

	fmt.Print(f.GetFooter())
}
