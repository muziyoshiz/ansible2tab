package main

import (
	"bufio"
	"fmt"
	"github.com/muziyoshiz/ansible2table/parser"
	"os"
	"github.com/muziyoshiz/ansible2table/formatter"
)

func main() {
	hosts := make([]string, 0, 10)
	results := make(map[string]parser.Result)

	parse := parser.Parser()

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		res, ok := parse(line)
		if ok {
			hosts = append(hosts, res.Host)
			results[res.Host] = res
		}
	}

	res, ok := parse(parser.EOF)
	if ok {
		hosts = append(hosts, res.Host)
		results[res.Host] = res
	}

	// Print TSV
	format := formatter.TsvFormatter()
	for _, host := range hosts {
		fmt.Print(format(results[host]))
	}
}
