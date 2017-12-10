# ansible2tab

Convert output of ansible command to TSV/JSON/Markdown/Backlog

## Description

ansible2tab converts output of ansible command to post it to blogs/wikis easily.

## Usage

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
$ ansible -i hosts app -m shell -a "cat /var/log/foobar | wc -l" | ansible2tab --format backlog
|Host|Value|h
|app1|177|
|app2|84|
|app3|37|
```

## Install

If you are OSX user, you can use [Homebrew](https://brew.sh/):

```
$ brew tap muziyoshiz/ansible2tab
$ brew install ansible2tab
```

If you are in another platform, you can download binary from [release page](https://github.com/muziyoshiz/ansible2tab/releases) and place it in `$PATH` directory.

Or you can use go get (you need to use go1.7 or later),

```
$ go get -u github.com/muziyoshiz/ansible2tab
```

## Develop

### Setup

```
$ go get github.com/mitchellh/gox
$ brew tap tcnksm/ghr
$ brew install ghr
```

### Build (TODO: Makefile)

```
$ gox -os "darwin linux windows" -arch "386 amd64" -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}/{{.Dir}}" -ldflags "-X main.revision=$(git rev-parse --short HEAD)"
$ for f in `ls pkg`; { zip -j pkg/${f}.zip pkg/${f}/ansible2tab* && rm -rf pkg/${f}; }
```

### Release (TODO: Makefile)

```
$ export GITHUB_TOKEN="....."
$ ghr -u muziyoshiz v<version> pkg/
```

## License

[MIT](https://github.com/muziyoshiz/ansible2tab/blob/master/LICENCE)

## Author

[Masahiro Yoshizawa](https://github.com/muziyoshiz)
