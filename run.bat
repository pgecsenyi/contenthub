@ECHO OFF

SET SAVED_GOPATH=%GOPATH%
SET GOPATH=%CD%

IF /i "%1" == "build" (
	go build -o bin/main.exe src/main.go
) ELSE IF /i "%1" == "fmt" (
	gofmt -w src
	gofmt -w src/configuration
	gofmt -w src/server
) ELSE IF /i "%1" == "lint" (
	golint src
	golint src/configuration
	golint src/server
) ELSE (
	echo No valid action provided.
)

SET GOPATH=%SAVED_GOPATH%
