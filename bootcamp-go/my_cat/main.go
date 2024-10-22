package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args[1:]

	for _, val := range args {
		data, err := os.ReadFile(val)

		if err != nil {
			Print("my_cat: " + val + ": No such file or directory")
			ap.PutRune('\n')
		} else {
			Print(string(data))
		}
	}
}

func Print(s string) {
	for i := 0; i < len(s); i++ {
		ap.PutRune(rune(s[i]))
	}
}
