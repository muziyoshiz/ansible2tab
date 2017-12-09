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

type JsonFormatter struct {
	defaultFormatter
	trailingLine bool
}

func (self *JsonFormatter) GetHeader() string {
	return "{"
}

func (self *JsonFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, "\\n")
	if self.trailingLine {
		return fmt.Sprintf(",\"%s\":\"%s\"", result.Host, values)
	} else {
		self.trailingLine = true
		return fmt.Sprintf("\"%s\":\"%s\"", result.Host, values)
	}
}

func (self *JsonFormatter) GetFooter() string {
	return "}"
}
