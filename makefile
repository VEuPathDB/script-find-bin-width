VERSION = $(shell git describe --tags --abbrev | sed 's/v\(.\+\)/\1/')

.PHONY: default
default:
	@echo "What are you doing?"

.PHONY: build
build: test
	@CGO_ENABLED=0 go build -o build/find-bin-width ./bin/find-bin-width/main.go

.PHONY: release
release: build
	@cd build/ && tar -czf fbw-$(VERSION).tar.gz find-bin-width

.PHONY: test
test: input-tests

.PHONY: input-tests
input-tests:
	@./input-test.sh
