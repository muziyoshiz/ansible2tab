package formatter

import (
	"github.com/muziyoshiz/ansible2table/parser"
	"strings"
	"fmt"
)

type defaultFormatter interface {
	GetHeader() string
	Format(result parser.Result) string
	GetFooter() string
}

type TsvFormatter struct {
	defaultFormatter
}

func (self *TsvFormatter) GetHeader() string {
	return ""
}

func (self *TsvFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, " ")
	return fmt.Sprintf("%s\t%s\n", result.Host, values)
}

func (self *TsvFormatter) GetFooter() string {
	return ""
}
