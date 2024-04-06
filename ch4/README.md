## BF interpreter

An interpreter for an esoteric programming language known as [brainfuck](https://esolangs.org/wiki/Brainfuck) often simply abbreviated as "BF". The language is best known for its almost trivial execution semantics \(so writing an interpreter for it is simple\) with only 8 possible commands, yet this is also precisely what makes it very difficult to write non-trivial programs in BF. BF is [Turing-complete](https://en.wikipedia.org/wiki/Turing_completeness) which means it is theoretically possible to implement any purely computational algorithm as a BF program.

This code example demonstrates the following Go concepts presented in the book:

- String manipulation
- Slices and maps
- More control flow: variations of the `for` loop in Go
- The `switch` construct and blank switches

### Developing

Follow the instructions in the project README and the last line of output after running `make` should be:

```text
Hello World!
```

The program interprets a Hello World program written in BF and displays the output of the BF program to STDOUT.

Customizing the BF program or input passed to the interpreter via `make` is not supported but feel free to replace the program and input passed to the interpreter by modifying the source code directly and re-running `make`.
