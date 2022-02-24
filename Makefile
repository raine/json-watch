.PHONY: build
build:
	go build -o bin/json-watch .

.PHONY: run
run:
	go run .

.PHONY: test
test: build
	./test/bats/bin/bats test/test.bats

.PHONY: go-test
go-test:
	richgo test

.PHONY: go-test-watch
go-test-watch:
	fd .go | entr -c richgo test

.PHONY: test-all
test-all: build go-test test
