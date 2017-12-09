# ansible2tab


## Develop

### Setup

```
$ go get github.com/motemen/gobump
$ go get github.com/mitchellh/gox
$ brew tap tcnksm/ghr
$ brew install ghr
```

### Build

```
$ gobump (major|minor|patch|set <version>)
$ gox -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}" -ldflags "-X main.revision=$(git rev-parse --short HEAD)"
```

### Release

```
$ export GITHUB_TOKEN="....."
$ ghr v0.0.1 pkg/
```
