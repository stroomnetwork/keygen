SHELL=/bin/bash

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

LDFLAGS := -X '$(PACKAGE_NAME)/version.Commit=$(shell git describe --always --tags --dirty)' \
	-X '$(PACKAGE_NAME)/version.CommitHash=$(shell git rev-parse HEAD)' \
	-X '$(PACKAGE_NAME)/version.GoVersion=$(shell go version)' \

.PHONY: all
all: help

## Build:

.PHONY: test
test: ## Run unit tests
	go test -v ./...


.PHONY: build
build: test ## Build the project
	go build -v -o ./build/keygen -installsuffix cgo -ldflags "$(LDFLAGS)" ./cmd/keygen

.PHONY: clean
clean: ## Remove build files and caches.
	rm -rf build
	go clean -i -r -cache -testcache

## Help:
.PHONY: help
help: ## Show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z0-9_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
