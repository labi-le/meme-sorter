.PHONY: build
.DEFAULT_GOAL := build

build-release:
	go build -ldflags "-s" -a -v -o build/package/memesorter-release cmd/main.go

build:
	go build -v -o build/package/memesorter-debug cmd/main.go


