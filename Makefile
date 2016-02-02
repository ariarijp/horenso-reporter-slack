BIN=horenso-reporter-slack

all: clean test

clean:
	rm -rf build/$(BIN)
	go clean

deps:
	go get -d -t -v ./...
	go get github.com/Songmu/horenso
	go get github.com/bluele/slack

test: deps
	go get github.com/stretchr/testify/assert
	go get github.com/pierrre/gotestcover
	gotestcover -v -covermode=count -coverprofile=.profile.cov ./...

.PHONY: all clean deps test
