GO_FILES      = $(shell find . -path ./vendor -prune -o -type f -name "*.go" -print)
IMPORT_PATH   = $(shell pwd | sed "s|^$(GOPATH)/src/||g")
LDFLAGS       = -w -X $(IMPORT_PATH)/pkg/version.PreRelease=$(PRE_RELEASE)

build: clean
	@go build -ldflags '$(LDFLAGS)'

clean:
	rm -f jukfs

test:
	go test $(shell go list ./... | grep -v /vendor/)

lint:
	@golint $(GO_FILES) || true

fmt:
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all clean test lint fmt
