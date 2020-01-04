GO := go
srcs := $(shell find . -path ./vendor -prune -o -name '*.go' | grep -v 'vendor')
PKGS ?= $(shell go list ./...)
PKG_FILES ?= *.go
GO_MINOR_VERSION := $(word 2,$(subst ., ,$(GO_VERSION)))
SHOULD_LINT := true

RACE=-race
GOTEST=go test -v $(RACE)
GOLINT=golint
GOVET=go vet
GOFMT=gofmt
FMT_LOG=fmt.log
LINT_LOG=lint.log


.PHONY: build 
build:
	@echo "Building the go module"
	go build ./...

.PHONY: fmt
fmt:
	$(GOFMT) -e -s -l -w $(srcs)

.PHONY: lint
lint:
ifdef SHOULD_LINT
	@rm -rf $(LINT_LOG)
	@rm -rf $(FMT_LOG)
	@echo "gofmt the files..."
	$(GOFMT) -e -s -l -w $(srcs) > $(FMT_LOG)
	@[ ! -s "$(FMT_LOG)" ] || (echo "Go Fmt Failures, run 'make fmt'" | cat - $(FMT_LOG) && false)
	@echo "Installing test dependencies for vet..."
	@go test -i $(PKGS)
	@echo "Checking vet..."
	$(GOVET) $(PKGS)
	@echo "Checking lint..."
	@$(foreach dir,$(PKGS),golint $(dir) 2>&1 | tee -a $(LINT_LOG);)
	@echo "Checking for unresolved FIXMEs..."
	@git grep -i fixme | grep -v -e vendor -e Makefile | tee -a $(LINT_LOG)
	@[ ! -s "$(LINT_LOG)" ] || (echo "Lint Failures" | cat - $(LINT_LOG) && false)
else
	@echo "Skipping linters on" $(GO_VERSION)
endif

.PHONY: test
test:
	$(GOTEST) $(PKGS)

.PHONY: cover
cover:
	go test -cover -coverprofile cover.out -race -v $(PKGS)

.PHONY: coveralls
coveralls:
	goveralls -service=travis-ci || echo "Coveralls failed"

