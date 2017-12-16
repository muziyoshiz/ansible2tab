package parser_test

import (
	"strings"
	"testing"

	"github.com/muziyoshiz/ansible2tab/parser"
)

// Works with one host
func TestParseWithOneHost(t *testing.T) {
	str := `app1 | SUCCESS | rc=0 >>
177
`
	lines := strings.Split(str, "\n")
	if len(lines) != 3 {
		t.Fatalf("Failed to create test data: length = %d", len(lines))
	}

	parse := parser.Parser()

	res, ok := parse(lines[0])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[1])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[2])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(parser.EOF)
	if res.Host != "app1" || len(res.Values) != 1 || res.Values[0] != "177" || res.Succeeded != true || res.Rc != 0 {
		t.Fatalf("Expected Result; got %+v", res)
	}
	if !ok {
		t.Fatalf("Expected true; got %t", ok)
	}
}

// Works with more than one host
func TestParseWithMoreThanOneHost(t *testing.T) {
	str := `app1 | SUCCESS | rc=0 >>
177

app2 | SUCCESS | rc=0 >>
84
`
	lines := strings.Split(str, "\n")
	if len(lines) != 6 {
		t.Fatalf("Failed to create test data: length = %d", len(lines))
	}

	parse := parser.Parser()

	res, ok := parse(lines[0])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[1])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[2])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[3])
	if res.Host != "app1" || len(res.Values) != 1 || res.Values[0] != "177" || res.Succeeded != true || res.Rc != 0 {
		t.Fatalf("Expected Result; got %+v", res)
	}
	if !ok {
		t.Fatalf("Expected true; got %t", ok)
	}

	res, ok = parse(lines[4])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[5])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(parser.EOF)
	if res.Host != "app2" || len(res.Values) != 1 || res.Values[0] != "84" || res.Succeeded != true || res.Rc != 0 {
		t.Fatalf("Expected Result; got %+v", res)
	}
	if !ok {
		t.Fatalf("Expected true; got %t", ok)
	}
}

// Return nothing with no host
func TestParseWithNoHost(t *testing.T) {
	str := `177
`
	lines := strings.Split(str, "\n")
	if len(lines) != 2 {
		t.Fatalf("Failed to create test data: length = %d", len(lines))
	}

	parse := parser.Parser()

	res, ok := parse(lines[0])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[1])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(parser.EOF)
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}
}

// Multiple line
func TestParseWithMultipleLine(t *testing.T) {
	str := `app1 | SUCCESS | rc=0 >>
177
ABC
`
	lines := strings.Split(str, "\n")
	if len(lines) != 4 {
		t.Fatalf("Failed to create test data: length = %d", len(lines))
	}

	parse := parser.Parser()

	res, ok := parse(lines[0])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[1])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[2])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(parser.EOF)
	if res.Host != "app1" || len(res.Values) != 2 || res.Values[0] != "177" || res.Values[1] != "ABC" || res.Succeeded != true || res.Rc != 0 {
		t.Fatalf("Expected Result; got %+v", res)
	}
	if !ok {
		t.Fatalf("Expected true; got %t", ok)
	}
}

// Header and no value (only with linefeed)
func TestParseWithNoLine(t *testing.T) {
	str := `app1 | SUCCESS | rc=0 >>
`
	lines := strings.Split(str, "\n")
	if len(lines) != 2 {
		t.Fatalf("Failed to create test data: length = %d", len(lines))
	}

	parse := parser.Parser()

	res, ok := parse(lines[0])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(lines[1])
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(parser.EOF)
	if res.Host != "app1" || len(res.Values) != 0 || res.Succeeded != true || res.Rc != 0 {
		t.Fatalf("Expected Result; got %+v", res)
	}
	if !ok {
		t.Fatalf("Expected true; got %t", ok)
	}
}

// Header without linefeed
func TestParseWithoutLinefeed(t *testing.T) {
	line := "app1 | SUCCESS | rc=0 >>"

	parse := parser.Parser()

	res, ok := parse(line)
	if res.Host != "" || len(res.Values) != 0 || res.Succeeded != false || res.Rc != 0 {
		t.Fatalf("Expected nil object; got %+v", res)
	}
	if ok {
		t.Fatalf("Expected false; got %t", ok)
	}

	res, ok = parse(parser.EOF)
	if res.Host != "app1" || len(res.Values) != 0 || res.Succeeded != true || res.Rc != 0 {
		t.Fatalf("Expected Result; got %+v", res)
	}
	if !ok {
		t.Fatalf("Expected true; got %t", ok)
	}
}
