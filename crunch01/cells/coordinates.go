package cells 

import "github.com/alem-platform/ap"
// creating letter sequence on x-axis
func Letters(width int) {
	var symbol int = 65
	var mul int = 1
	for i := 0; i < width; i++ {
		for j := 0; j < 7 - mul + 1; j++ {
			ap.PutRune(' ')
		}
		for j := 0; j < mul; j++ {
			ap.PutRune(rune(symbol))
		}
		symbol++
		if symbol > 'Z' {
			symbol = 65
			mul++
		}
	}
	ap.PutRune('\n')
}
// creating digit sequence on y-axis
func Digits(stage int, number int) {
	var spaces int = 3
	if stage == 1 {
		var copy int = number
		for copy > 0 {
			spaces--
			copy /= 10
		}
		PutNumber(number)
		for i := 0; i < spaces; i++ {
			ap.PutRune(' ')
		}
	} else {
		for i := 0; i < spaces; i++ {
			ap.PutRune(' ')
		}
	}
}