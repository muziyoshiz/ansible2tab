package formatter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/muziyoshiz/ansible2tab/parser"
)

// Formatter provides basic features for formatting ansible outputs.
type Formatter interface {
	GetHeader() string
	Format(result parser.Result) string
	GetFooter() string
}

// TSVFormatter is a stateless formatter for TSV.
type TSVFormatter struct {
	Formatter
}

// GetHeader returns blank string. ansible2tab prints TSV without any header.
func (*TSVFormatter) GetHeader() string {
	return ""
}

// Format returns TSV with linefeed
func (*TSVFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, " ")
	return fmt.Sprintf("%s\t%s\n", result.Host, values)
}

// GetFooter returns blank string. TSV does not have any footer.
func (*TSVFormatter) GetFooter() string {
	return ""
}

// JSONFormatter is a stateful formatter for JSON.
type JSONFormatter struct {
	Formatter
	trailingLine bool
}

// GetHeader returns JSON object header
func (*JSONFormatter) GetHeader() string {
	return "{"
}

// Format returns a JSON pair.
// If it's not first pair, one separating "," is also returned.
func (f *JSONFormatter) Format(result parser.Result) string {
	jsHost, _ := json.Marshal(result.Host)
	jsValues, _ := json.Marshal(strings.Join(result.Values, "\n"))

	if f.trailingLine {
		return fmt.Sprintf(",%s:%s", jsHost, jsValues)
	}

	f.trailingLine = true
	return fmt.Sprintf("%s:%s", jsHost, jsValues)
}

// GetFooter returns JSON object footer with linefeed
func (*JSONFormatter) GetFooter() string {
	return "}\n"
}

// MarkdownFormatter is a stateless formatter for Markdown table.
type MarkdownFormatter struct {
	Formatter
}

// GetHeader returns a header row of Markdown table
func (*MarkdownFormatter) GetHeader() string {
	return `|Host|Value|
|---|---|
`
}

// Format returns a row of Markdown table
func (*MarkdownFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, " ")
	return fmt.Sprintf("|%s|%s|\n", result.Host, values)
}

// GetFooter returns blank string. Markdown table does not have any footer.
func (*MarkdownFormatter) GetFooter() string {
	return ""
}

// MarkdownCodeFormatter is a stateful formatter for Markdown code block.
type MarkdownCodeFormatter struct {
	Formatter
	trailingLine bool
}

// GetHeader returns blank string. Group of Code block does not have any header.
func (*MarkdownCodeFormatter) GetHeader() string {
	return ""
}

// Format returns a Markdown code block.
// If it's not first code block, one separating linefeed is also returned.
func (f *MarkdownCodeFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, "\n")

	if f.trailingLine {
		// We can not escape backquote inside backquotes
		return fmt.Sprintf("\n## %s\n\n```\n%s\n```\n", result.Host, values)
	}

	f.trailingLine = true
	return fmt.Sprintf("## %s\n\n```\n%s\n```\n", result.Host, values)
}

// GetFooter returns blank string. Group of Code block does not have any footer.
func (*MarkdownCodeFormatter) GetFooter() string {
	return ""
}

// BacklogFormatter is a stateless formatter for Backlog-format table.
type BacklogFormatter struct {
	Formatter
}

// GetHeader returns a header row of Backlog-format table
func (*BacklogFormatter) GetHeader() string {
	return "|Host|Value|h\n"
}

// Format returns a row of Backlog-format table
func (*BacklogFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, " ")
	return fmt.Sprintf("|%s|%s|\n", result.Host, values)
}

// GetFooter returns blank string. Backlog-format table does not have any footer.
func (*BacklogFormatter) GetFooter() string {
	return ""
}

// BacklogCodeFormatter is a stateful formatter for Backlog-format code block.
type BacklogCodeFormatter struct {
	Formatter
	trailingLine bool
}

// GetHeader returns blank string. Group of Code block does not have any header.
func (*BacklogCodeFormatter) GetHeader() string {
	return ""
}

// Format returns a Backlog-style code block.
// If it's not first code block, one separating linefeed is also returned.
func (f *BacklogCodeFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, "\n")
	if f.trailingLine {
		// We can not escape backquote inside backquotes
		return fmt.Sprintf("\n** %s\n\n{code}\n%s\n{/code}\n", result.Host, values)
	}

	f.trailingLine = true
	return fmt.Sprintf("** %s\n\n{code}\n%s\n{/code}\n", result.Host, values)
}

// GetFooter returns blank string. Group of Code block does not have any footer.
func (*BacklogCodeFormatter) GetFooter() string {
	return ""
}
