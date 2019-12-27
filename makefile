#!/bin/bash

export PKGS=$(shell go list ./... | grep -v vendor/)

all:
	build test 

build:
	build-server-for-linux
build-server-for-linux:
	@echo "Building server for linux..."
	@export GOOS=linux GOARCH=amd64
	@go build -o ./bin/linux/server server.go
	@echo "Done."

build-server-for-windows:
	@echo "Building server for windows..."
	@GOOS=windows GOARCH=amd64
	@go build -o ./bin/windows/server.exe server.go
	@echo "Done."

build-server-for-mac:
	@echo "Building server for mac..."
	@GOOS=mac GOARCH=amd64
	@go build -o ./bin/mac/server server.go
	@echo "Done."



build-client-for-linux:
	@echo "Building client for linux..."
	@export GOOS=linux GOARCH=amd64
	@go build -o ./bin/linux/client ./client/client.go
	@echo "Done."

build-client-for-windows:
	@echo "Building client for windows..."
	@GOOS=windows GOARCH=amd64
	@go build -o ./bin/windows/client.exe ./client/client.go
	@echo "Done."

build-client-for-mac:
	@echo "Building client for mac..."
	@GOOS=mac GOARCH=amd64
	@go build -o ./bin/mac/client ./client/client.go
	@echo "Done."

