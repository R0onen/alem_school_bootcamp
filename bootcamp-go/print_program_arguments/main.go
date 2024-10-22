package main

import "os"

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		os.Stdout.WriteString(args[i] + "\n")
	}
}
