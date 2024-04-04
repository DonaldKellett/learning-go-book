package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBrainfuckHelloWorld(t *testing.T) {
	b, e := os.ReadFile("progs/hello.b")
	check(e)
	prog := string(b)
	expected := "Hello World!"
	actual, err := brainfuck(prog, "")
	assert.Nil(t, err)
	assert.Equal(t, expected, actual, "should return \"Hello World!\"")
}

func TestBrainfuckCatProgram(t *testing.T) {
	prog, input, expected := ",[.,]", "", ""
	actual, err := brainfuck(prog, input)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual, "should return the empty string")
	input, expected = "Golang is awesome <3", "Golang is awesome <3"
	actual, err = brainfuck(prog, input)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual, "should return \"Golang is awesome <3\"")
}

func TestBrainfuckUTMSimulation(t *testing.T) {
	b, e := os.ReadFile("progs/utm.b")
	check(e)
	prog, input := string(b), "b1b1bbb1c1c11111d"
	expected := "1c11111\n"
	actual, err := brainfuck(prog, input)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual, "should work for Daniel B Cristofani's UTM simulation")
}

func TestBrainfuckMismatchedSquareBrackets(t *testing.T) {
	prog, expectedErrorString := "]", "mismatched ']'"
	actual, err := brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
	prog, expectedErrorString = "][", "mismatched ']'"
	actual, err = brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
	prog, expectedErrorString = "[", "mismatched '['"
	actual, err = brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
	prog, expectedErrorString = "[[]", "mismatched '['"
	actual, err = brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
}

func TestBrainfuckInvalidTapePointerDereference(t *testing.T) {
	prog, expectedErrorString := "<.", "invalid tape pointer dereference when < 0"
	actual, err := brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
	prog, expectedErrorString = "<,", "invalid tape pointer dereference when < 0"
	actual, err = brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
	prog, expectedErrorString = "<[]", "invalid tape pointer dereference when < 0"
	actual, err = brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
	prog, expectedErrorString = "+[>+]", "invalid tape pointer dereference when >= 30000"
	actual, err = brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
	prog, expectedErrorString = "-[>-]", "invalid tape pointer dereference when >= 30000"
	actual, err = brainfuck(prog, "")
	assert.EqualError(t, err, expectedErrorString)
	assert.Empty(t, actual)
}

func TestBrainfuckTapePointerOutOfBoundsNoDereference(t *testing.T) {
	prog, expected := "<x", ""
	actual, err := brainfuck(prog, "")
	assert.Nil(t, err)
	assert.Equal(t, expected, actual, "should return the empty string")
}
