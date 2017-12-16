package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// formulaTmpl is temaplate for homebrew formula
var formulaTmpl = `require "formula"
class {{ .Name | Title }} < Formula
  desc "Convert output of ansible command to TSV/JSON/Markdown/Backlog"
  homepage "https://github.com/muziyoshiz/{{ .Name }}"
  version '{{ .Version }}'
  url "https://github.com/muziyoshiz/{{ .Name }}/releases/download/{{ .Version }}/{{ .Name }}_{{ .Version }}_darwin_amd64.zip"
  sha256 "{{ .Sha256 }}"

  def install
    bin.install '{{ .Name }}'
  end
end
`

// c.f. https://github.com/tcnksm/ghr/blob/master/release/main.go
func main() {
	os.Exit(_main())
}

func _main() int {
	if len(os.Args) != 4 {
		log.Println("Usage: go run main.go NAME VERSION FILE")
		return 0
	}

	name := os.Args[1]
	version := os.Args[2]
	file := os.Args[3]

	file, err := filepath.Abs(file)
	if err != nil {
		log.Println(err)
		return 1
	}

	f, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return 1
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return 1
	}
	checkSum := sha256.Sum256(buf)

	tmpl, err := template.New("formula").Funcs(template.FuncMap{
		"Title": strings.Title,
	}).Parse(formulaTmpl)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Execute(os.Stdout, struct {
		Name, Version, Sha256 string
	}{
		Name:    name,
		Version: version,
		Sha256:  fmt.Sprintf("%x", checkSum),
	}); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
