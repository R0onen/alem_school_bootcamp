package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func ConvertStringToASCII(input []rune) {
	for i, r := range input {
		asciiValue := int(r)
		ascii100 := asciiValue / 100
		ascii10 := (asciiValue / 10) % 10
		ascii1 := asciiValue % 10

		if asciiValue >= 100 {
			ap.PutRune(rune('0' + ascii100))
		}
		if asciiValue >= 10 {
			ap.PutRune(rune('0' + ascii10))
		}
		ap.PutRune(rune('0' + ascii1))

		if i < len(input)-1 {
			ap.PutRune(' ')
		}
	}
}

func main() {
	var N int
	fmt.Scanf("%d", &N)

	input := make([]rune, N)
	for i := 0; i < len(input); i++ {
		var r rune
		fmt.Scanf("%c", &r)
		input[i] = r
	}
	ConvertStringToASCII(input)
}
