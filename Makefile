GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=internal


GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: help

## Build:
build: ## Build your project and put the output binary in out/bin/
	echo "Compilando..."
	mkdir ./bin
	$(GOCMD) build -o ./bin/$(BINARY_NAME) ./internal/

clean: ## Remove build related file
	rm -fr ./bin

## Test:
test: ## Run the tests of the project
	echo "Testing..."

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