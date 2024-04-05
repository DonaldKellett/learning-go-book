N=94

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
	go build
.PHONY:build

run: build
	./learning-go-book-ch3 -n ${N}
.PHONY:run

clean:
	rm -vf ./learning-go-book-ch3
.PHONY:clean
