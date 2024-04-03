package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	expected, actual := "Hello, world!", hello()
	assert.Equal(t, expected, actual, "Function hello() should return \"Hello, world!\"")
}
