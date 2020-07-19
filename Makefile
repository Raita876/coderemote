VERSION := v1.0.0
PACKAGE_NAME := coderemote

BUILD_GOOS := darwin
BUILD_GOARCH := amd64

.PHONY: build
build:
	GOOS=$(BUILD_GOOS) \
	GOARCH=$(BUILD_GOARCH) \
	go build \
		-o ./bin/$(PACKAGE_NAME) \
		-ldflags "-X main.version=$(VERSION) -X main.name=$(PACKAGE_NAME)" .

.PHONY: install
install: build
	chmod 755 ./bin/$(PACKAGE_NAME) && mv ./bin/$(PACKAGE_NAME) /usr/local/bin/

.PHONY: tag
tag:
	git tag $(VERSION)
	git push origin $(VERSION)
