SHELL := /bin/bash

GO_FILES := $(shell find terminal-dance -name '*.go')
GO_MODULE_DIRS := $(shell find terminal-dance -path '*/go/go.mod' -print | sed 's#/go.mod##')
RUST_MANIFESTS := $(shell find terminal-dance -path '*/rust/Cargo.toml' -print)

.PHONY: ci frontend-style python-style gofmt golangci go-test cargo-fmt cargo-clippy cargo-test

ci: frontend-style python-style gofmt golangci go-test cargo-fmt cargo-clippy cargo-test

frontend-style:
	@echo 'Running Prettier and ESLint checks...'
	npm run prettier
	npm run eslint

python-style:
	@echo 'Running Black and Ruff checks...'
	npm run black
	npm run ruff

gofmt:
	@echo 'Checking gofmt formatting...'
	@if [ -n "$(GO_FILES)" ]; then \
		unformatted=$$(gofmt -l $(GO_FILES)); \
		if [ -n "$$unformatted" ]; then \
			echo "The following Go files need gofmt:"; \
			echo "$$unformatted"; \
			exit 1; \
		fi; \
	else \
		echo 'No Go files found under terminal-dance.'; \
	fi

golangci:
	@echo 'Running golangci-lint...'
	@if [ -n "$(GO_MODULE_DIRS)" ]; then \
		for dir in $(GO_MODULE_DIRS); do \
			echo "golangci-lint in $$dir"; \
			( cd $$dir && golangci-lint run ); \
		done; \
	else \
		echo 'No Go modules found under terminal-dance.'; \
	fi

go-test:
	@echo 'Running go test...'
	@if [ -n "$(GO_MODULE_DIRS)" ]; then \
		for dir in $(GO_MODULE_DIRS); do \
			echo "go test ./... in $$dir"; \
			( cd $$dir && go test ./... ); \
		done; \
	else \
		echo 'No Go modules found under terminal-dance.'; \
	fi

cargo-fmt:
	@echo 'Checking cargo fmt formatting...'
	@if [ -n "$(RUST_MANIFESTS)" ]; then \
		for manifest in $(RUST_MANIFESTS); do \
			echo "cargo fmt --manifest-path $$manifest -- --check"; \
			cargo fmt --manifest-path $$manifest -- --check; \
		done; \
	else \
		echo 'No Rust workspaces found under terminal-dance.'; \
	fi

cargo-clippy:
	@echo 'Running cargo clippy...'
	@if [ -n "$(RUST_MANIFESTS)" ]; then \
		for manifest in $(RUST_MANIFESTS); do \
			echo "cargo clippy --manifest-path $$manifest --all-targets --all-features -- -D warnings"; \
			cargo clippy --manifest-path $$manifest --all-targets --all-features -- -D warnings; \
		done; \
	else \
		echo 'No Rust workspaces found under terminal-dance.'; \
	fi

cargo-test:
	@echo 'Running cargo test...'
	@if [ -n "$(RUST_MANIFESTS)" ]; then \
		for manifest in $(RUST_MANIFESTS); do \
			echo "cargo test --manifest-path $$manifest"; \
			cargo test --manifest-path $$manifest; \
		done; \
	else \
		echo 'No Rust workspaces found under terminal-dance.'; \
	fi
