package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	if len(os.Args) < 2 {
		PrintLn("Usage: go run main.go <brainfuck_code>")
		return
	}

	brainfuckCode := os.Args[1]
	interpretBrainfuck(brainfuckCode)
}

func interpretBrainfuck(code string) {
	mem := make([]byte, 30000)
	ptr := 0
	codeIndex := 0
	loopStack := []int{}

	for codeIndex < len(code) {
		switch code[codeIndex] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '.':
			ap.PutRune(rune(mem[ptr]))
		case ',':
			if len(os.Args) > 1 {

				input := os.Args[1]

				if len(input) > 0 {
					mem[ptr] = input[0]
				} else {
					mem[ptr] = 0
				}

			} else {
				mem[ptr] = 0
			}
		case '[':
			if mem[ptr] == 0 {
				loopDepth := 1
				for loopDepth > 0 {
					codeIndex++
					if code[codeIndex] == '[' {
						loopDepth++
					} else if code[codeIndex] == ']' {
						loopDepth--
					}
				}
			} else {
				// Push current code index onto the stack
				loopStack = append(loopStack, codeIndex)
			}
		case ']':
			if mem[ptr] != 0 {
				// Jump back to the matching '['
				codeIndex = loopStack[len(loopStack)-1]
			} else {
				// Pop from loop stack
				loopStack = loopStack[:len(loopStack)-1]
			}
		}
		codeIndex++
	}
}

func Print(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func PrintLn(s string) {
	Print(s + "\n")
}
