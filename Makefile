SHELL := /usr/bin/env bash -o pipefail

pre:
	# install buf CLI, protoc-gen-buf-breaking, and protoc-gen-buf-lint binaries, bash completion, fish completion, and zsh completion
	PREFIX="/usr/local"; \
	VERSION="1.32.1"; \
	curl -sSL "https://github.com/bufbuild/buf/releases/download/v$$VERSION/buf-$(shell uname -s)-$(shell uname -m).tar.gz" | \
	sudo tar -xvzf - -C "$$PREFIX" --strip-components 1

	# install protoc plugin for Go
	go install google.golang.org/protobuf/cmd/protoc-gen-go

start-server:
	go run cmd/server/main.go

start-client:
	go run cmd/client/main.go
