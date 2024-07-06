BINARY_NAME := IsNumeric

MODULE_PATH := ex09
RM          := rm -rf


all: init fmt build

init:
	@if [ ! -f go.mod ]; then \
		go mod init ${MODULE_PATH}; \
	fi

build: ${SRC}
	go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

fmt:
	go fmt *.go
	go fmt vendor/piscine/*.go

clean:
	go clean
	${RM} go.mod

fclean: clean
	${RM} ${BINARY_NAME}

re: fclean all

.PHONY: init, build, run, fmt, clean, fclean, re
