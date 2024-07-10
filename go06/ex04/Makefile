BINARY_NAME := ztail

MODULE_PATH := ex03
TEST_DIR		:= test
TEST_NAME		:= ztail_test.go
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
	go fmt test/*.go

test: init
	go build -o $(TEST_DIR)/$(BINARY_NAME) main.go
	go test -v $(TEST_DIR)/$(TEST_NAME) --short

clean:
	go clean
	${RM} go.mod $(TEST_DIR)/test.test

fclean: clean
	${RM} ${BINARY_NAME} $(TEST_DIR)/${BINARY_NAME}

re: fclean all

.PHONY: init, build, run, fmt, clean, fclean, re
