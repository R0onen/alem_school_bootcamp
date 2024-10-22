package bootcamp

import (
	"bootcamp/firststruct"

	"github.com/alem-platform/ap"
)

func PrintUserInfo(u firststruct.User) {
	HasPassword := "no"
	if u.PasswordReliability() != "undefined" {
		HasPassword = "yes"
	}
	PrintWords("Name: ")
	PrintWords(u.Name)
	ap.PutRune('\n')
	PrintWords("HasPassword: ")
	PrintWords(HasPassword)
	ap.PutRune('\n')
}

func PrintWords(word string) {
	for _, letter := range word {
		ap.PutRune(rune(letter))
	}
}
