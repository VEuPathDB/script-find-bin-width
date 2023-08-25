.PHONY: default
default:
	@echo "What are you doing?"

.PHONY: build
build:
	@CGO_ENABLED=0 go build -o build/find-bin-width ./bin/find-bin-width/main.go

.PHONY: test
test: input-tests

.PHONY: input-tests
input-tests:
	@./input-test.sh
