package main

import "github.com/alem-platform/ap"

func main() {
	letterslower := 'a'
	lettersupper := 'A'
	for i := 0; i < 26; i++ {
		ap.PutRune(letterslower)
		letterslower++
		ap.PutRune(lettersupper)
		lettersupper++
	}
	ap.PutRune('\n')
}
