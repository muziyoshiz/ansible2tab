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
	formatOpt = flag.StringP("format", "f","tsv", `output format "tsv", "js", "md" or "md-code"`)
)

func main() {
	flag.Parse()

	var f formatter.Formatter
	switch *formatOpt {
	case "tsv":
		f = &formatter.TsvFormatter{}
	case "js", "json":
		f = &formatter.JsonFormatter{}
	case "md", "markdown":
		f = &formatter.MarkdownFormatter{}
	case "md-code", "markdown-code":
		f = &formatter.MarkdownCodeFormatter{}
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
