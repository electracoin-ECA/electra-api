NAME = electra-api

ifeq ($(OS), Windows_NT)
	BINARY_NAME = ${NAME}.exe
else
	BINARY_NAME = ${NAME}
endif

install:
	go get -u github.com/tools/godep
	go get -u github.com/stretchr/testify
	go get -u golang.org/x/lint/golint
	godep restore

lint:
	golint ./main.go
	golint ./src/...

start:
	go build && "./${BINARY_NAME}"

test:
	make lint
	go test -v ./...
