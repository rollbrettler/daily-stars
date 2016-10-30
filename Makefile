PWD=$(shell pwd)
BUILD_DIR?=$(shell pwd)/build
GOX_OS?=linux darwin windows solaris freebsd netbsd openbsd

test:
	go test -race ./...

crosscompile:
	go get github.com/mitchellh/gox
	mkdir -p ${BUILD_DIR}/bin
	gox -output="${BUILD_DIR}/bin/{{.Dir}}-{{.OS}}-{{.Arch}}" -os="${GOX_OS}" ${GOX_FLAGS} ./...
