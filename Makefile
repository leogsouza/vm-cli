GO_CMD=go
GO_BUILD=${GO_CMD} build
BINARY_NAME=vmcli
BINARY_DIR=bin

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO_BUILD} -o ${BINARY_DIR}/${BINARY_NAME}
build-windows:
	${GO_BUILD} -o ${BINARY_DIR}/${BINARY_NAME}.exe

build: build-windows

run: build
	@./${BINARY_DIR}/${BINARY_NAME}

test:
	@go test -cover -race -v ./...