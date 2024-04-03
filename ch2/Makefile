N=93

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
.PHONY:vet

ci-vet: fmt-check
	go vet ./...
.PHONY:ci-vet

test: ci-vet
	go test
.PHONY:test

build: vet
	go build fib.go
.PHONY:build

run: build
	./fib -n ${N}
.PHONY:run

clean:
	rm -vf ./fib
.PHONY:clean
