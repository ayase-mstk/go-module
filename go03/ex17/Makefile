BINARY_NAME := AtoiBase

MODULE_PATH := ex17
TEST_DIR		:= test
TEST_NAME		:= atoibase_test.go
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

test: init
	go test -v $(TEST_DIR)/$(TEST_NAME) --short

clean:
	go clean
	${RM} go.mod

fclean: clean
	${RM} ${BINARY_NAME}

re: fclean all

.PHONY: init, build, run, fmt, clean, fclean, re
