package parser

import (
	"regexp"
	"strconv"
)

// EOF is a special string to teach the end of ansible output to the formatters
const EOF string = "EOF"

// Result means a result against one host
type Result struct {
	Host      string
	Values    []string
	Succeeded bool
	Rc        int
}

// Parser returns a closure to parse ansible output
func Parser() func(string) (Result, bool) {
	var header = regexp.MustCompile(`([^ ]+) \| (SUCCESS|FAILED) \| rc=(\d+) >>`)
	var host, prevHost = "", ""
	var values, prevValues = make([]string, 0, 10), make([]string, 0, 10)
	var succeeded, prevSucceeded = false, false
	var rc, prevRc = 0, 0

	return func(line string) (Result, bool) {
		if line == EOF {
			if host == "" {
				return Result{}, false
			}

			// If the last value is blank, remove the last value
			num := len(values)
			if num > 0 {
				lastValueLen := len(values[num-1])
				if lastValueLen == 0 {
					values = values[:num-1]
				}
			}

			return Result{host, values, succeeded, rc}, true
		}

		if group := header.FindSubmatch([]byte(line)); len(group) != 0 {
			prevHost, prevValues, prevSucceeded, prevRc = host, values, succeeded, rc

			host = string(group[1])
			succeeded = bool(string(group[2]) == "SUCCESS")
			rc, _ = strconv.Atoi(string(group[3]))

			values = make([]string, 0, 10)

			if prevHost == "" {
				return Result{}, false
			}

			// If the last value is blank, remove the last value
			num := len(prevValues)
			if num > 0 {
				lastValueLen := len(prevValues[num-1])
				if lastValueLen == 0 {
					prevValues = prevValues[:num-1]
				}
			}

			return Result{prevHost, prevValues, prevSucceeded, prevRc}, true
		} else if host != "" {
			values = append(values, line)
		}
		return Result{}, false
	}
}
