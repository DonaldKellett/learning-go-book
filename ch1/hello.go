package main

import "fmt"

func hello() string {
	return "Hello, world!"
}

func main() {
	var greeting = hello()
	fmt.Println(greeting)
}
