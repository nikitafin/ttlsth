GOCMD  = go
GOTEST = $(GOCMD) test
GOVET  = $(GOCMD) vet

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor

all: help

ECHO_CMD := @echo
UTILS  := $(wildcard ./cmd/*/*)

## Build:
build: ## Build your project and put the output binary in out/bin/
	@echo Building...
	@mkdir -p out/bin
	@$(foreach var,$(UTILS), $(GOCMD) build -o out/bin $(var);)

clean: ## Remove build related file
	@echo Cleaning...
	@rm -fr ./bin
	@rm -fr ./out

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)