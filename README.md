# ansible2tab


## Develop

### Setup

```
$ go get github.com/mitchellh/gox
$ brew tap tcnksm/ghr
$ brew install ghr
```

### Build

```
$ gox -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}"
```

### Release

```
$ export GITHUB_TOKEN="....."
$ ghr v0.0.1 pkg/
```
