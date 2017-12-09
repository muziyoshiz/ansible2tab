package main

import (
	"bufio"
	"fmt"
	"github.com/muziyoshiz/ansible2tab/parser"
	"os"
	"github.com/muziyoshiz/ansible2tab/formatter"
)

func main() {
	parse := parser.Parser()
	//f := formatter.TsvFormatter{}
	f := formatter.MarkdownCodeFormatter{}

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
