package main

import "github.com/alem-platform/ap"

func main() {
	letters := 'z'
	for i := 0; i < 26; i++ {
		ap.PutRune(letters)
		letters--
	}
	ap.PutRune('\n')
}
