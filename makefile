VERSION = $(shell git describe --tags --abbrev=0 | sed 's/v\(.\+\)/\1/')

.PHONY: default
default:
	@echo "What are you doing?"


##
#
#  BUILD TASKS
#
##

.PHONY: build
build: build-windows build-darwin build-linux

.PHONY: build-windows
build-windows:
	@echo "Compiling for windows"
	@mkdir -p build/windows
	@env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/windows/find-bin-width.exe ./bin/find-bin-width/main.go

.PHONY: build-darwin
build-darwin:
	@echo "Compiling for darwin"
	@mkdir -p build/darwin
	@env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/darwin/find-bin-width ./bin/find-bin-width/main.go

.PHONY: build-linux
build-linux:
	@echo "Compiling for linux"
	@mkdir -p build/linux
	@env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/linux/find-bin-width ./bin/find-bin-width/main.go


##
#
#  RELEASE TASKS
#
##


.PHONY: release
release: release-windows release-darwin release-linux

.PHONY: release-windows
release-windows: build-windows
	@cd build/windows \
		&& zip -9 fbw-windows-$(VERSION).zip find-bin-width.exe \
		&& mv fbw-windows-$(VERSION).zip ..

.PHONY: release-darwin
release-darwin: build-darwin
	@cd build/darwin \
		&& zip -9 fbw-darwin-$(VERSION).zip find-bin-width \
		&& mv fbw-darwin-$(VERSION).zip ..

.PHONY: release-linux
release-linux: build-linux
	@cd build/linux \
		&& zip -9 fbw-linux-$(VERSION).zip find-bin-width \
		&& mv fbw-linux-$(VERSION).zip ..

##
#
#  TEST TASKS
#
##


.PHONY: test
test: input-tests

.PHONY: input-tests
input-tests: build-linux
	@./input-test.sh build/linux/find-bin-width

