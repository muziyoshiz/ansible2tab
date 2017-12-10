# ansible2tab

Convert output of ansible command to TSV/JSON/Markdown/Backlog

## Description

ansible2tab converts output of ansible command to post it to blogs/wikis easily.

## usage

Without ansible2tab:

```
$ ansible -i hosts app -m shell -a "cat /var/log/foobar | wc -l"
app1 | SUCCESS | rc=0 >>
177

app2 | SUCCESS | rc=0 >>
84

app3 | SUCCESS | rc=0 >>
37
```

With ansible2tab:

```
$ ansible -i hosts app -m shell -a "cat /var/log/foobar | wc -l" | ansible2tab
app1	177
app2	84
app3	37
$ ansible -i hosts app -m shell -a "cat /var/log/foobar | wc -l" | ansible2tab --format json
{"app1":"177","app2":"84","app3":"37"}
$ ansible -i hosts app -m shell -a "cat /var/log/foobar | wc -l" | ansible2tab --format markdown
|Host|Value|
|---|---|
|app1|177|
|app2|84|
|app3|37|
```

## Install

```
$ go get github.com/muziyoshiz/ansible2tab
```

## Develop

### Setup

```
$ go get github.com/mitchellh/gox
$ brew tap tcnksm/ghr
$ brew install ghr
```

### Build

```
$ gox -os "darwin linux windows" -arch "386 amd64" -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}/{{.Dir}}" -ldflags "-X main.revision=$(git rev-parse --short HEAD)"
$ for f in `ls pkg`; { zip -j pkg/${f}.zip pkg/${f}/ansible2tab* && rm -rf pkg/${f}; }
```

### Release

```
$ export GITHUB_TOKEN="....."
$ ghr v<version> pkg/
```

## License

[MIT](https://github.com/muziyoshiz/ansible2tab/blob/master/LICENCE)

## Author

[Masahiro Yoshizawa](https://github.com/muziyoshiz)
