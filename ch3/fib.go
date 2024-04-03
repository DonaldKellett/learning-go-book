package main

import (
	"flag"
	"fmt"
)

var (
	n uint64
)

func fib(n uint64) []uint64 {
	ns := make([]uint64, 0, n)
	if n == 0 {
		return ns
	}
	ns = append(ns, 0)
	if n == 1 {
		return ns
	}
	ns = append(ns, 1)
	var a, b uint64 = 0, 1
	for n > 2 {
		a, b = b, a+b
		ns = append(ns, b)
		n -= 1
	}
	return ns
}

func main() {
	flag.Uint64Var(&n, "n", 0, "An integer 0 <= n < 95")
	flag.Parse()
	if n < 0 || n >= 95 {
		s := fmt.Sprintf("Expected input integer n = %d to be within the bounds 0 <= n < 95", n)
		panic(s)
	}
	fmt.Println("The first", n, "Fibonacci numbers are:", fib(n))
}
