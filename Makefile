
PROJECTNAME := $(shell basename "$(PWD)")
BASE := $(shell pwd)
BIN := $(BASE)/bin
FILES := ./src/*.go

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean $(FILES)
GOTEST=$(GOCMD) test ./... -cover

.DEFAULT_GOAL := help
.PHONY: help

# -----------------------------------------------------------------------------
# General
# -----------------------------------------------------------------------------

get: build ## Builds and run Benchhog with get requests.
	$(BIN)/$(PROJECTNAME) get google.com

post: build ## Builds and run Benchhog with post requests.
	$(BIN)/$(PROJECTNAME) post google.com

build: ## Builds Benchhog to bin directory.
	$(GOBUILD) -o $(BIN)/$(PROJECTNAME) $(FILES)

test: ## Run unit tests
	$(GOTEST)

clean: ## Remove artifacts
	$(GOCLEAN)
	rm -fr $(BIN) 2> /dev/null


# -----------------------------------------------------------------------------
# Git
# -----------------------------------------------------------------------------

new-release: ## Release new version of software.
	@./scripts/new-release.sh

# -----------------------------------------------------------------------------
# Other
# -----------------------------------------------------------------------------

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
