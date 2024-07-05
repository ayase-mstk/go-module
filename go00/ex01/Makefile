BINARY_NAME := PrintReverseAlphabet

MODULE_PATH := ex01
RM = rm -rf


all: init fmt build

init:
	go mod init ${MODULE_PATH}

build: ${SRC}
	go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

fmt:
	go fmt *.go

test:
	go test -v ./... -short

clean:
	go clean
	${RM} go.mod

fclean: clean
	${RM} ${BINARY_NAME}

re: fclean all

.PHONY: init, build, run, fmt, test, clean, fclean, re
