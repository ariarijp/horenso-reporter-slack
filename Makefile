BIN=horenso-reporter-slack

all: clean test build

build: deps
	go build -o build/$(BIN)

clean:
	rm -rf build/$(BIN)
	go clean

deps:
	go get -d -t -v ./...
	go get github.com/Songmu/horenso
	go get github.com/bluele/slack

run: build
	./build/$(BIN)

test: deps
	go get github.com/stretchr/testify/assert
	go get github.com/pierrre/gotestcover
	gotestcover -v -covermode=count -coverprofile=.profile.cov ./...

.PHONY: all build clean deps run test
