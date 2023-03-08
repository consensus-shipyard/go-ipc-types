targetdir := bin

all: test samples

.PHONY: help
help:
	@echo 'Usage: make [target]...'
	@echo ''
	@echo 'Generic targets:'
	@echo '  all (default)    - Build and test all'
	@echo '  test             - Run all tests'
	@echo '  lint             - Run code quality checks'
	@echo '  generate         - Generate dependent files'

.PHONY: test
test:
	go test -v -shuffle=on -count=1 -race -timeout 20m ./...

.PHONY: test_cov
test_cov:
	go test -coverprofile coverage.out -v -count=1 -timeout 20m ./...

.PHONY: format
format:
	gofmt -w -s .
	goimports -w -local "github.com/filecoin-project,github.com/consensus-shipyard/go-ipc-types" .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: vulncheck
vulncheck:
	govulncheck -v ./...

.PHONY: generate
generate:
	gofmt -w -s .
	go run gen/gen.go
