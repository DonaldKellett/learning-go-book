package main

import (
	"fmt"
	"os"
)

func brainfuck(code string, input string) (string, error) {
	tokens := []rune(code)
	bytesIn := []byte(input)
	bytesInIdx := 0
	bytesOut := []byte{}
	parensMap := map[int]int{}
	parensMapRev := map[int]int{}
	parensStack := []int{}
	for i, r := range tokens {
		switch r {
		case '[':
			parensStack = append(parensStack, i)
		case ']':
			if len(parensStack) == 0 {
				return "", fmt.Errorf("mismatched ']'")
			}
			j := len(parensStack) - 1
			parensMap[parensStack[j]] = i
			parensMapRev[i] = parensStack[j]
			parensStack = parensStack[:j]
		}
	}
	if len(parensStack) > 0 {
		return "", fmt.Errorf("mismatched '['")
	}
	tape := make([]byte, 30_000)
	tapeIdx := 0
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case '>':
			tapeIdx++
		case '<':
			tapeIdx--
		case '+', '-', '.', ',', '[', ']':
			switch {
			case tapeIdx < 0:
				return "", fmt.Errorf("invalid tape pointer dereference when < 0")
			case tapeIdx >= 30_000:
				return "", fmt.Errorf("invalid tape pointer dereference when >= 30000")
			default:
				switch tokens[i] {
				case '+':
					tape[tapeIdx]++
				case '-':
					tape[tapeIdx]--
				case '.':
					bytesOut = append(bytesOut, tape[tapeIdx])
				case ',':
					if bytesInIdx >= len(bytesIn) {
						tape[tapeIdx] = 0
					} else {
						tape[tapeIdx] = bytesIn[bytesInIdx]
						bytesInIdx++
					}
				case '[':
					if tape[tapeIdx] == 0 {
						i = parensMap[i]
					}
				case ']':
					if tape[tapeIdx] != 0 {
						i = parensMapRev[i]
					}
				}
			}
		}
	}
	output := string(bytesOut)
	return output, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	b, e := os.ReadFile("progs/hello.b")
	check(e)
	prog := string(b)
	output, err := brainfuck(prog, "")
	check(err)
	fmt.Println(output)
}
