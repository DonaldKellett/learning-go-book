package main

import "fmt"

func hello() string {
	return "Hello, world!"
}

func main() {
	s := hello()
	fmt.Println(s)
}
