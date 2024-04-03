# learning-go-book

My adaptation of the code examples and exercises in [Learning Go](https://learning-go-book.dev/) by Jon Bodner

## Developing

Fork and clone this repository, then navigate to the project root and make it your working directory.

The code examples for each chapter is self-contained with [Make](https://www.gnu.org/software/make/) as the build system. For example, to run the code examples for chapter 1, enter the `ch1/` subdirectory and invoke the following command:

```bash
make
```

To run unit tests instead \(used in CI\):

```bash
make test
```

## License

[MIT](./LICENSE)
