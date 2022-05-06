GOCMD  = go
GOTEST = $(GOCMD) test
GOVET  = $(GOCMD) vet

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0  )

.PHONY: all help

all: help

# WIX_SOURCE  := $(wildcard ./wix/*/*)
WIX_DIR = ./wix
WIX_ARRAYS_STRINGS := $(dir $(wildcard $(WIX_DIR)/arrays-strings/*/))

## Build:
build: ## no main, run only tests
	@mkdir -p out/bin
	@$(foreach var,$(WIX_ARRAYS_STRINGS), $(GOTEST) $(var);)

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