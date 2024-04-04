# learning-go-book

My adaptation of the code examples and exercises in [Learning Go](https://learning-go-book.dev/) by Jon Bodner

## Developing

### Prerequisites

Ensure [Make](https://www.gnu.org/software/make/) is installed, then download and install a sufficiently recent version of [Go](https://go.dev/) \(>= 1.22.1\), set your `GOPATH` accordingly, add `$GOPATH/bin` to your `PATH` and install the shadow linter:

```bash
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
```

### Running the code examples

Fork and clone this repository, then navigate to the project root and make it your working directory.

The code examples for each chapter is self-contained with Make as the build system. For example, to run the code examples for chapter 1, enter the `ch1/` subdirectory and invoke the following command:

```bash
make
```

To run unit tests instead \(used in CI\):

```bash
make test
```

## License

[MIT](./LICENSE)
