VERSION = $(shell git describe --tags --abbrev=0 | sed 's/v\(.\+\)/\1/')

C_BLUE = "\\033[94m"
C_NONE = "\\033[0m"

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
	@echo "$(C_BLUE)Compiling for windows$(C_NONE)"
	@mkdir -p build/windows
	@env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/windows/find-bin-width.exe ./bin/find-bin-width/main.go

.PHONY: build-darwin
build-darwin:
	@echo "$(C_BLUE)Compiling for darwin$(C_NONE)"
	@mkdir -p build/darwin
	@env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/darwin/find-bin-width ./bin/find-bin-width/main.go

.PHONY: build-linux
build-linux:
	@echo "$(C_BLUE)Compiling for linux$(C_NONE)"
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
	@echo "$(C_BLUE)Packaging for windows$(C_NONE)"
	@cd build/windows \
		&& zip -9 fbw-windows-$(VERSION).zip find-bin-width.exe \
		&& mv fbw-windows-$(VERSION).zip ..

.PHONY: release-darwin
release-darwin: build-darwin
	@echo "$(C_BLUE)Packaging for darwin$(C_NONE)"
	@cd build/darwin \
		&& zip -9 fbw-darwin-$(VERSION).zip find-bin-width \
		&& mv fbw-darwin-$(VERSION).zip ..

.PHONY: release-linux
release-linux: build-linux
	@echo "$(C_BLUE)Packaging for linux$(C_NONE)"
	@cd build/linux \
		&& zip -9 fbw-linux-$(VERSION).zip find-bin-width \
		&& mv fbw-linux-$(VERSION).zip ..

##
#
#  TEST TASKS
#
##


.PHONY: test
test: unit-test input-test

.PHONY: input-test
input-test: build-linux
	@echo "$(C_BLUE)Running input tests$(C_NONE)"
	@./input-test.sh build/linux/find-bin-width

.PHONY: unit-test
unit-test:
	@echo "$(C_BLUE)Running unit tests$(C_NONE)"
	@go test ./...

