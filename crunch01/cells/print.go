package cells

import "github.com/alem-platform/ap"

// function for printing words
func PrintWords(word string) {
	for _, letter := range word {
		ap.PutRune(rune(letter))
	}
	ap.PutRune('\n')
}
