package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	var expected = "Hello, world!"
	var actual = hello()
	assert.Equal(t, expected, actual, "Function hello() should return \"Hello, world!\"")
}
