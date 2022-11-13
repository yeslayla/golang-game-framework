PROJECTNAME="Golang Game Framework"
PROJECT_BIN="golang-game-framework"
VERSION="DEV"
BUILD_NUMBER:=$(shell git rev-parse --short HEAD)

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/build
GOFILES=$(wildcard *.go)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

go-full-build: go-clean go-get go-build-static

go-build:
	@echo "  >  Building binary..."
	@mkdir -p $(GOBIN)
	@GOOS=linux go build -ldflags "-X github.com/yeslayla/$(PROJECT_BIN)/core.Version=$(VERSION) -X github.com/yeslayla/$(PROJECT_BIN)/core.Build=$(BUILD_NUMBER)" -o $(GOBIN)/$(PROJECT_BIN) $(GOFILES)
	@chmod 755 $(GOBIN)/$(PROJECT_BIN)

go-build-static:
	@echo "  >  Building static binary..."
	@mkdir -p $(GOBIN)
	@CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags "-s -w -X github.com/yeslayla/$(PROJECT_BIN)/core.Version=$(VERSION) -X github.com/yeslayla/$(PROJECT_BIN)/core.Build=$(BUILD_NUMBER)" -o $(GOBIN)/$(PROJECT_BIN) $(GOFILES)
	@chmod 755 $(GOBIN)/$(PROJECT_BIN)

go-build-windows:
	@echo "  >  Building static binary..."
	@mkdir -p $(GOBIN)
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w -X github.com/yeslayla/$(PROJECT_BIN)/core.Version=$(VERSION) -X github.com/yeslayla/$(PROJECT_BIN)/core.Build=$(BUILD_NUMBER) -H windowsgui" -o $(GOBIN)/$(PROJECT_BIN).exe $(GOFILES)
	@chmod 755 $(GOBIN)/$(PROJECT_BIN)

go-generate:
	@echo "  >  Generating dependency files..."
	@go generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get $(get)

go-install:
	@echo "  >  Running go install..."
	@go install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@go clean

go-test: clean
	@echo "  >  Running tests..."
	@go test -coverprofile=coverage.out ./*/

go-run:
	@echo "  >  Running ${PROJECTNAME}"
	@-(cd $(GOBIN); ./$(PROJECT_BIN))

## install: Download and install dependencies
install: go-get
	@sudo mv ./build/$(PROJECT_BIN) /usr/bin/$(PROJECT_BIN)
	@sudo chmod +x /usr/bin/$(PROJECT_BIN)

# clean: Runs go clean
clean: go-clean

## full-build: cleans project, installs dependencies, and builds project
full-build: go-full-build

## build-windows: Builds project staticly for Windows
build-windows: go-build-windows

## build: Runs go build
build: go-build

## package: Builds lambda zip
package: go-full-build
	@echo "  >  Zipping package..."
	@cd $(GOBIN) && zip $(PROJECTNAME).zip $(PROJECTNAME)

## clean: Runs go clean
clean:
	@rm -rf build

## run: full-builds and executes project binary
run: go-build go-run

## test: Run unit tests
test: go-test

## help: Displays help text for make commands
.DEFAULT_GOAL := help
all: help
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'