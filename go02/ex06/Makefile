BINARY_NAME := IsPrime

MODULE_PATH := ex06
SRC_DIR		:= ./vendor/piscine
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
	go fmt $(SRC_DIR)/*.go

clean:
	go clean
	${RM} go.mod

fclean: clean
	${RM} ${BINARY_NAME}

re: fclean all

.PHONY: init, build, run, fmt, clean, fclean, re
