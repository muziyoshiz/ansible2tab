package main

import (
	"bufio"
	"fmt"
	"github.com/muziyoshiz/ansible2table/parser"
	"os"
	"strings"
)

func main() {
	hosts := make([]string, 0, 10)
	values := make(map[string][]string)

	parse := parser.Parser()

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		res, ok := parse(line)
		if ok {
			hosts = append(hosts, res.Host)
			values[res.Host] = res.Values
		}
	}

	res, ok := parse(parser.EOF)
	if ok {
		hosts = append(hosts, res.Host)
		values[res.Host] = res.Values
	}

	// Print TSV
	for _, host := range hosts {
		value := strings.Join(values[host], " ")
		fmt.Printf("%s\t%s\n", host, value)
	}
}
