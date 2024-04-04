.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

fmt-check:
	test -z "$$(gofmt -l .)"
.PHONY:fmt-check

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
	shadow ./...
.PHONY:vet

ci-vet: fmt-check
	go vet ./...
	shadow ./...
.PHONY:ci-vet

test: ci-vet
	go test
.PHONY:test

build: vet
	go build brainfuck.go
.PHONY:build

run: build
	./brainfuck
.PHONY:run

clean:
	rm -vf ./brainfuck
.PHONY:clean
