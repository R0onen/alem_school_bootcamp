package bootcamp

import "github.com/alem-platform/ap"

func PrintAlphabetLn() {
	letter := 'a'

	for i := 0; i < 26; i++ {
		ap.PutRune(letter)
		letter++
	}
	ap.PutRune('\n')
}

func main() {
	PrintAlphabetLn()
}
