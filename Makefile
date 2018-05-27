
# define build version
VERSION ?= 1.0.0
BUILD_COMMIT := `git rev-parse HEAD`
BUILD_TIME := `date`
GO_VERSION := `go version`

# global go tools
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=gofmt

# go build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X 'main.BuildCommit=$(BUILD_COMMIT)' \
-X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GoVersion=$(GO_VERSION)'"

# binary name
BINARY := goplayer
PLATFORMS := windows linux darwin
OS = $(word 1, $@)

GO_FILES=$(shell find . -type f -name "*.go" -not -path "./pkg/*")

release: windows linux darwin

$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARC=amd64 $(GOBUILD) $(LDFLAGS) -o release/$(BINARY)-$(VERSION)-$(OS)-amd64

gentest:
	gotests -all -excl main -w $(GO_FILES)

test:
	$(GOTEST) -v -cover=true ./...

fmt:
	$(GOFMT) -l -w $(GO_FILES)

check:
	@test -z $(shell gofmt -l $(GO_FILES) | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${GO_FILES}

clean:
	$(GOCLEAN)
	rm -rf release

.PHONY: release $(PLATFORMS)
