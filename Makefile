BINARY   := geelark-cli
MODULE   := github.com/geelark-tech/geelark-cli
VERSION  := $(shell git describe --tags --always 2>/dev/null || echo dev)
DATE     := $(shell date +%Y-%m-%d)
TIMESTAMP := $(shell date +%Y%m%d%H%M)
LDFLAGS  := -s -w -X $(MODULE)/internal/build.Version=$(VERSION) -X $(MODULE)/internal/build.Date=$(DATE)
PREFIX   ?= /usr/local
DIST     := dist
OUTDIR   := $(DIST)/$(TIMESTAMP)

.PHONY: build vet test install uninstall clean release dist

RELEASE_VERSION :=

release:
ifndef RELEASE_VERSION
	$(error Usage: make release RELEASE_VERSION=x.y.z)
endif
	@echo "=== Release $(RELEASE_VERSION) ==="
	git tag "$(RELEASE_VERSION)"
	@echo "✔ Tagged $(RELEASE_VERSION)"
	$(MAKE) dist
	@echo "✔ Built dist"
	./npm/publish.sh "$(RELEASE_VERSION)"
	@echo "✔ Published to npm"
	git checkout main
	@echo "✔ Switched back to main"
	@echo "=== Release $(RELEASE_VERSION) complete ==="

build:
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY) .

vet:
	go vet ./...

test:
	go test -race -count=1 ./...

install: build
	install -d $(PREFIX)/bin
	install -m755 $(BINARY) $(PREFIX)/bin/$(BINARY)
	@echo "OK: $(PREFIX)/bin/$(BINARY) ($(VERSION))"

uninstall:
	rm -f $(PREFIX)/bin/$(BINARY)

clean:
	rm -f $(BINARY)
	rm -rf $(DIST)

dist: dist-darwin-arm64 dist-darwin-amd64 dist-linux-amd64 dist-windows-amd64

dist-darwin-arm64:
	@mkdir -p $(OUTDIR)
	GOOS=darwin GOARCH=arm64 go build -trimpath -ldflags "$(LDFLAGS)" -o $(OUTDIR)/$(BINARY)-darwin-arm64 .
	@echo "✔ $(OUTDIR)/$(BINARY)-darwin-arm64"

dist-darwin-amd64:
	@mkdir -p $(OUTDIR)
	GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o $(OUTDIR)/$(BINARY)-darwin-amd64 .
	@echo "✔ $(OUTDIR)/$(BINARY)-darwin-amd64"

dist-linux-amd64:
	@mkdir -p $(OUTDIR)
	GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o $(OUTDIR)/$(BINARY)-linux-amd64 .
	@echo "✔ $(OUTDIR)/$(BINARY)-linux-amd64"

dist-windows-amd64:
	@mkdir -p $(OUTDIR)
	GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o $(OUTDIR)/$(BINARY)-windows-amd64.exe .
	@echo "✔ $(OUTDIR)/$(BINARY)-windows-amd64.exe"
