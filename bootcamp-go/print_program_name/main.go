package main

import (
	"os"
)

func main() {
	programName := os.Args[0]

	lastSlash := 0
	for i := 0; i < len(programName); i++ {
		if programName[i] == '/' {
			lastSlash = i
		}
	}

	programName = programName[lastSlash+1:]

	os.Stdout.WriteString(programName + "\n")
}
