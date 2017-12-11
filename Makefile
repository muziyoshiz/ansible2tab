VERSION = $(shell grep 'const version' version.go | sed -E 's/.*"(.+)"$$/\1/')
PACKAGES = "./parser ./formatter"
FORMULA = "../homebrew-ansible2tab/ansible2tab.rb"

all: build package upload

build:
	gox -os "darwin linux windows" -arch "386 amd64" -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}/{{.Dir}}" -ldflags "-X main.revision=$(git rev-parse --short HEAD)"

install:
	go install

package: build
	@for f in $(shell ls pkg); { zip -j pkg/$$f.zip pkg/$$f/ansible2tab* && rm -rf pkg/$$f; }

release:
	ghr -u ${GITHUB_USER} v${VERSION} pkg/

# It works but shows "make: Nothing to be done for `brew'."
brew:
	$(eval URL := "https:\/\/github.com\/muziyoshiz\/ansible2tab\/releases\/download\/v${VERSION}\/ansible2tab_darwin_amd64.zip")
	$(eval CHECKSUM := $(shell shasum -a 256 pkg/ansible2tab_darwin_amd64.zip | awk '{print $$1;}'))
	$(shell sed -i '' -E 's/url .*$$/url ${URL}/' ${FORMULA})
	$(shell sed -i '' -E 's/version .*$$/version \"v${VERSION}\"/' ${FORMULA})
	$(shell sed -i '' -E 's/sha256 .*$$/sha256 \"v${CHECKSUM}\"/' ${FORMULA})

test-all: vet lint test

test:
	go test -v -parallel=4 ${PACKAGES}

test-race:
	go test -v -race ${PACKAGES}

vet:
	go vet ${PACKAGES}

lint:
	@go get github.com/golang/lint/golint
	go list ./... | grep -v vendor | xargs -n1 golint 

clean:
	-rm -rf pkg

.PHONY: all build install package release brew test-all test test-race vet lint clean
