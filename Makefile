SHELL := /usr/bin/env bash -o pipefail

COMPOSE_FILE_PATH := build/docker-compose.yaml

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

build-images:
	@echo "# building docker images"
	docker-compose -f ${COMPOSE_FILE_PATH} build --force-rm --pull user_service gateway_service

start-services:
	@echo "# starting user(server) and gateway(client) service"
	cd build && docker-compose up -d

stop-services:
	@echo "# stopping user(server) and gateway(client) service"
	cd build && docker-compose down