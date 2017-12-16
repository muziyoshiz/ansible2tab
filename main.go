package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/muziyoshiz/ansible2tab/formatter"
	"github.com/muziyoshiz/ansible2tab/parser"
	flag "github.com/spf13/pflag"
)

var (
	formatOpt  = flag.StringP("format", "f", "tsv", `output format "tsv", "js", "md", "md-code", "blg" or "blg-code"`)
	versionOpt = flag.BoolP("version", "v", false, `version`)
	helpOpt    = flag.BoolP("help", "h", false, `help`)
)

func main() {
	os.Exit(_main())
}

func _main() int {
	flag.Parse()

	if *helpOpt {
		flag.PrintDefaults()
		return 0
	}

	if *versionOpt {
		fmt.Printf("ansible2tab version: %s (commit: %s)\n", version, commit)
		return 0
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
		flag.PrintDefaults()
		return 1
	}

	parse := parser.Parser()

	fmt.Print(f.GetHeader())

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		res, ok := parse(line)
		if ok {
			fmt.Print(f.Format(res))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input: ", err)
	}

	res, ok := parse(parser.EOF)
	if ok {
		fmt.Print(f.Format(res))
	}

	fmt.Print(f.GetFooter())
	return 0
}
