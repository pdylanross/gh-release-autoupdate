ifneq (,$(wildcard ./.env))
    include .env
    export
endif


help: # Show this help.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: default
default: help

.PHONY: fix-deps
fix-deps: # Run dependency maintenance commands
	go mod tidy
	go vet ./...

.PHONY: fix-check
fix-check: fix-deps # Run linters and autofix issues
	golangci-lint run --fix

.PHONY: check
check: # Run linters, no autofix
	golangci-lint run

.PHONY: unit-test
unit-test: # Run unit tests
	go clean -testcache
	go test ./...

.PHONY: pre-commit
pre-commit: fix-check unit-test # Run all standard cleanups before a commit
