# GoFunAndChallenges — common dev tasks.
# Run `make help` for the full list.

.DEFAULT_GOAL := help
SHELL         := /bin/bash

.PHONY: help
help: ## Show available targets.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: tidy
tidy: ## Sync go.mod / go.sum.
	go mod tidy

.PHONY: fmt
fmt: ## Format all Go source.
	gofmt -s -w .

.PHONY: vet
vet: ## Run go vet across all packages.
	go vet ./...

.PHONY: build
build: ## Build every package (does not produce binaries).
	go build ./...

.PHONY: test
test: ## Run all tests.
	go test ./...

.PHONY: test-race
test-race: ## Run all tests with the race detector.
	go test -race ./...

.PHONY: cover
cover: ## Produce a coverage profile and HTML report.
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Open coverage.html in a browser."

.PHONY: bench
bench: ## Run benchmarks across all packages.
	go test -bench=. -benchmem ./...

.PHONY: lint
lint: ## Run golangci-lint (install: https://golangci-lint.run/).
	@command -v golangci-lint >/dev/null 2>&1 || { \
		echo "golangci-lint not installed. Install: https://golangci-lint.run/usage/install/"; \
		exit 1; \
	}
	golangci-lint run ./...

.PHONY: check
check: fmt vet test ## fmt + vet + test — the default sanity check.

.PHONY: hello
hello: ## Run lesson 01 — quickest sanity check.
	go run ./lessons/01-hello

.PHONY: clean
clean: ## Remove build & coverage artifacts.
	rm -f coverage.out coverage.html
	go clean ./...
