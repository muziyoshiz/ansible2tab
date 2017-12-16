NAME = ansible2tab
VERSION = $(shell grep 'const version' version.go | sed -E 's/.*"(.+)"$$/\1/')
COMMIT = $(shell git describe --always)
PACKAGES = $(shell go list ./...)
FORMULA = ../homebrew-$(NAME)/$(NAME).rb
EXTERNAL_TOOLS = github.com/mitchellh/gox github.com/tcnksm/ghr

default: test

# Install external tools for this project
bootstrap:
	@for tool in $(EXTERNAL_TOOLS) ; do \
		echo "Installing $$tool" ; \
    	go get $$tool; \
done

# Build binary on './bin' directory for local use
build:
	go build -ldflags "-X main.commit=$(COMMIT)" -o bin/$(NAME)

# Install binary on $GOPATH/bin directory
install:
	go install -ldflags "-X main.commit=$(COMMIT)"

# Build zip files on './pkg/dist/$(VERSION)' directory
package:
	@sh -c "'$(CURDIR)/scripts/package.sh' $(NAME) $(VERSION) $(COMMIT)"

# Overwrite ../homebrew-$(NAME)/$(NAME).rb
brew: package
	go run release/main.go $(NAME) $(VERSION) pkg/dist/$(VERSION)/$(NAME)_$(VERSION)_darwin_amd64.zip > ../homebrew-$(NAME)/$(NAME).rb

# Release zip files to GitHub
upload:
	ghr -u $(GITHUB_USER) $(VERSION) pkg/dist/$(VERSION)

test-all: vet lint test

test:
	go test -v -parallel=4 ${PACKAGES}

test-race:
	go test -v -race ${PACKAGES}

vet:
	go vet -composites=false ${PACKAGES}

lint:
	@go get github.com/golang/lint/golint
	go list ./... | grep -v vendor | xargs -n1 golint 

.PHONY: bootstrap build install package brew upload test-all test test-race vet lint
