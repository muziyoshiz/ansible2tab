package formatter_test

import (
	"github.com/muziyoshiz/ansible2tab/formatter"
	"testing"
	"github.com/muziyoshiz/ansible2tab/parser"
)

// Works with one value
func TestTsvFormatterWithOneValue(t *testing.T) {
	f := formatter.TsvFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "app1\t177\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with two values
func TestTsvFormatterWithTwoValues(t *testing.T) {
	f := formatter.TsvFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	values = append(values, "ABC")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "app1\t177 ABC\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with one value
func TestJsonFormatterWithOneValue(t *testing.T) {
	f := formatter.JsonFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "\"app1\":\"177\""

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with two values
func TestJsonFormatterWithTwoValues(t *testing.T) {
	f := formatter.JsonFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	values = append(values, "ABC")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "\"app1\":\"177\\nABC\""

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with one value
func TestJsonFormatterWithTwoHosts(t *testing.T) {
	f := formatter.JsonFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "\"app1\":\"177\""

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}

	values = make([]string, 0, 1)
	values = append(values, "84")
	actual = f.Format(parser.Result{"app2", values, true, 0})
	expected = ",\"app2\":\"84\""

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with one value
func TestMarkdownFormatterWithOneValue(t *testing.T) {
	f := formatter.MarkdownFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "|app1|177|\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with two values
func TestMarkdownFormatterWithTwoValues(t *testing.T) {
	f := formatter.MarkdownFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	values = append(values, "ABC")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "|app1|177 ABC|\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with one value
func TestMarkdownCodeFormatterWithOneValue(t *testing.T) {
	f := formatter.MarkdownCodeFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	// We can not escape backquote inside backquotes
	expected := "## app1\n\n```\n177\n```\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with two values
func TestMarkdownCodeFormatterWithTwoValues(t *testing.T) {
	f := formatter.MarkdownCodeFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	values = append(values, "ABC")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "## app1\n\n```\n177\nABC\n```\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with one value
func TestMarkdownCodeFormatterWithTwoHosts(t *testing.T) {
	f := formatter.MarkdownCodeFormatter{}

	values := make([]string, 0, 1)
	values = append(values, "177")
	actual := f.Format(parser.Result{"app1", values, true, 0})
	expected := "## app1\n\n```\n177\n```\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}

	values = make([]string, 0, 1)
	values = append(values, "84")
	actual = f.Format(parser.Result{"app2", values, true, 0})
	expected = "\n## app2\n\n```\n84\n```\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}
