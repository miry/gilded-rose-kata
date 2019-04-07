GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTOOL=$(GOCMD) tool
GOTEST=$(GOCMD) test -v -race
GOGET=$(GOCMD) get
GOLINT=golangci-lint
BINARY_NAME=bin/gilded_rose

all: test build

run: build
	./$(BINARY_NAME)

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) ./...

coverage:
	$(GOTEST) ./... -coverprofile=coverage.out
	$(GOTOOL) cover -html=coverage.out

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

lint:
	$(GOLINT) run
