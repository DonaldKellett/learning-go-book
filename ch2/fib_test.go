package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibBaseCases(t *testing.T) {
	var expected, actual uint64 = 0, fib(0)
	assert.Equal(t, expected, actual, "f(0) = 0")
	expected, actual = 1, fib(1)
	assert.Equal(t, expected, actual, "f(1) = 1")
}

func TestFibSmallNumbers(t *testing.T) {
	var expected, actual uint64 = 1, fib(2)
	assert.Equal(t, expected, actual, "f(2) = 1")
	expected, actual = 2, fib(3)
	assert.Equal(t, expected, actual, "f(3) = 2")
	expected, actual = 3, fib(4)
	assert.Equal(t, expected, actual, "f(4) = 3")
	expected, actual = 5, fib(5)
	assert.Equal(t, expected, actual, "f(5)= 5")
	expected, actual = 55, fib(10)
	assert.Equal(t, expected, actual, "f(10) = 55")
}

func TestFibLargeNumbers(t *testing.T) {
	var expected, actual uint64 = 832_040, fib(30)
	assert.Equal(t, expected, actual, "f(30) = 832040")
	expected, actual = 12_586_269_025, fib(50)
	assert.Equal(t, expected, actual, "f(50) = 12586269025")
	expected, actual = 12_200_160_415_121_876_738, fib(93)
	assert.Equal(t, expected, actual, "f(93) = 12200160415121876738")
}
