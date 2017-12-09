package formatter_test

import (
	"github.com/muziyoshiz/ansible2table/formatter"
	"testing"
	"github.com/muziyoshiz/ansible2table/parser"
)

// Works with one value
func TestTsvFormatterWithOneValue(t *testing.T) {
	format := formatter.TsvFormatter()

	values := make([]string, 0, 1)
	values = append(values, "177")
	actual := format(parser.Result{"app1", values, true, 0})
	expected := "app1\t177\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}

// Works with two values
func TestTsvFormatterWithTwoValues(t *testing.T) {
	format := formatter.TsvFormatter()

	values := make([]string, 0, 1)
	values = append(values, "177")
	values = append(values, "ABC")
	actual := format(parser.Result{"app1", values, true, 0})
	expected := "app1\t177 ABC\n"

	if actual != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, actual)
	}
}
