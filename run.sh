#!/bin/sh
SAVED_GOPATH=$GOPATH
export GOPATH=$(pwd)

if [ "$1" = "build" ]; then
    go build -o bin/main src/main.go
elif [ "$1" = "lint" ]; then
    golint src
	golint src/configuration
	golint src/server
elif [ "$1" = "fmt" ]; then
	gofmt -w src
	gofmt -w src/configuration
	gofmt -w src/server
else
	echo "No valid action provided."
fi

export GOPATH=$SAVED_GOPATH
