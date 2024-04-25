MAIN_PACKAGE_PATH := ./
BINARY_NAME := simple

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the application
.PHONY: build
build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
	pigeon -o=${MAIN_PACKAGE_PATH}internal/parser/pegparser/markdown.go ${MAIN_PACKAGE_PATH}internal/parser/pegparser/markdown.peg

## run: run the webserver
.PHONY: run-webserver
run-webserver: build
	/tmp/bin/${BINARY_NAME}	webserver

## run: run the repl
.PHONY: run-repl
run-repl: build
	/tmp/bin/${BINARY_NAME}	repl