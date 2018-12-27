#!make
#-----------------------------------------------------------------------------
# Copyright (C) Microsoft. All rights reserved.
# Licensed under the MIT license.
# See LICENSE.txt file in the project root for full license information.
#-----------------------------------------------------------------------------

GO_BIN ?= go
GO_LINT ?= golint
BINARY_NAME ?= cleanDeployment
IMAGE_NAME ?= cleanDeployment

.PHONY: all
all: dep build

.PHONY: build-docker
build-docker: 
	GOOS=linux GOARCH=amd64 $(GO_BIN) build -o bin/$(BINARY_NAME)
	docker build --rm -t ${IMAGE_NAME} .

.PHONY: build
build:
	$(GO_BIN) build -o bin/$(BINARY_NAME) .

.PHONY: mac-build
mac-build:
	GOOS=darwin GOARCH=amd64 $(GO_BIN) build -o bin/$(BINARY_NAME) .

.PHONY: linux-build
linux-build:
	GOOS=linux GOARCH=amd64 $(GO_BIN) build -o bin/$(BINARY_NAME) .

.PHONY: win-build
win-build:
	GOOS=windows GOARCH=amd64 $(GO_BIN) build -o bin/$(BINARY_NAME) .

.PHONY: lint
lint:
	$(GO_LINT) .

.PHONY: dep
dep:
	glide install --strip-vendor

.PHONY: clean
clean:
	rm -rf bin
	rm -rf vendor
	$(GO_BIN) clean

.DEFAULT_GOAL := build