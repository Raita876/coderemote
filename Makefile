VERSION := v0.1.0
PACKAGE_NAME := coderemote

.PHONY: build
build:
	go build \
		-o ./bin/$(PACKAGE_NAME) \
		-ldflags "-X main.version=$(VERSION) -X main.name=$(PACKAGE_NAME)" .
