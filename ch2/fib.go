package main

import (
	"flag"
	"fmt"
)

var (
	n uint64
)

func fib(n uint64) uint64 {
	if n < 2 {
		return n
	}
	var a, b uint64 = 1, 1
	for n > 2 {
		a, b = b, a+b
		n -= 1
	}
	return b
}

func main() {
	flag.Uint64Var(&n, "n", 0, "An integer 0 <= n < 94")
	flag.Parse()
	if n < 0 || n >= 94 {
		s := fmt.Sprintf("Expected input integer n = %d to be within the bounds 0 <= n < 94", n)
		panic(s)
	}
	s := fmt.Sprintf("f(%d) = %d", n, fib(n))
	fmt.Println(s)
}
