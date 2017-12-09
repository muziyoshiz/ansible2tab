package formatter

import (
	"github.com/muziyoshiz/ansible2tab/parser"
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
	return "}\n"
}

type MarkdownFormatter struct {
	defaultFormatter
}

func (self *MarkdownFormatter) GetHeader() string {
	return `|Host|Value|
|---|---|
`
}

func (self *MarkdownFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, " ")
	return fmt.Sprintf("|%s|%s|\n", result.Host, values)
}

func (self *MarkdownFormatter) GetFooter() string {
	return ""
}

type MarkdownCodeFormatter struct {
	defaultFormatter
	trailingLine bool
}

func (self *MarkdownCodeFormatter) GetHeader() string {
	return ""
}

func (self *MarkdownCodeFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, "\n")
	if self.trailingLine {
		// We can not escape backquote inside backquotes
		return fmt.Sprintf("\n## %s\n\n```\n%s\n```\n", result.Host, values)
	} else {
		self.trailingLine = true
		return fmt.Sprintf("## %s\n\n```\n%s\n```\n", result.Host, values)
	}
}

func (self *MarkdownCodeFormatter) GetFooter() string {
	return ""
}
