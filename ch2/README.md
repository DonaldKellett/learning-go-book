## Fibonacci number

A simple program that computes the nth Fibonacci number, demonstrating the following Go concepts presented in the book:

- Variable declaration and initialization
- Built-in integer types
- Variable assignment
- Basic control flow

### Developing

Follow the instructions in the project README and the last line of output after running `make` should be:

```text
f(93) = 12200160415121876738
```

By default, the program computes the 93rd Fibonacci number which is configurable. For example, to compute the 10th Fibonacci number instead, run:

```bash
make N=10
```

Which results in the following output:

```text
f(10) = 55
```
