BINARY_NAME=mdtool
VERSION=$(shell git describe --tags --always --abbrev=0 || echo "v0.1.0")
BUILD_FLAGS=-ldflags "-X main.version=${VERSION}"

.PHONY: build
build:
	go build ${BUILD_FLAGS} -o ${BINARY_NAME} ./cmd/mdtool

.PHONY: install
install:
	go install ${BUILD_FLAGS} ./cmd/mdtool

.PHONY: clean
clean:
	rm -f ${BINARY_NAME}
