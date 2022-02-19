.PHONY: build
build:
	go build -o bin/json-watch .

.PHONY: run
run:
	go run .

.PHONY: test
test: build
	./test/bats/bin/bats test/test.bats
