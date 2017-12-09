package formatter

import (
	"github.com/muziyoshiz/ansible2table/parser"
	"strings"
	"fmt"
)

func TsvFormatter() func(result parser.Result) string {
	return func(result parser.Result) string {
		values := strings.Join(result.Values, " ")
		return fmt.Sprintf("%s\t%s\n", result.Host, values)
	}
}
