BINARY = ela
BUILD_DIR = ".build"

define build_goarch
	@echo "Building server for ${1} ${2}."
	@env GOOS=${1} GOARCH=${2} go build -buildmode=c-shared -o django_lib/ela/ela.so django_lib/ela/ela.go
endef

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: clean
## clean: clean build directory
clean:
	@echo Cleaning ${BUILD_DIR}.
	@rm -rf ${BUILD_DIR}
	@echo Done cleaning ${BUILD_DIR}.

.PHONY: build
## build: build binary
build:
	$(call build_goarch,darwin,amd64)
	@echo Done mac build.
	$(call build_goarch,linux,amd64)
	@echo Done linux build.
