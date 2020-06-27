#!/bin/sh -ex
go clean -testcache
go test ./...
go fmt ./...
go vet ./...
VER=$(git tag | grep '^v' | tail -1)
REV=$(git rev-parse --short HEAD)
go build -ldflags "-s -w -X main.Version=$VER -X main.Revision=$REV"
