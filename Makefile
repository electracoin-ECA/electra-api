# Variables
BINARY_NAME=electra-api
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test

# Commands
all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
deps:
	godep save
run:
	$(GOBUILD)
	./$(BINARY_NAME)
test:
	$(GOTEST) -v ./...

# OS Compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
build-docker:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_NAME)" -v
