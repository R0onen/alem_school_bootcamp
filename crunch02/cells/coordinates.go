package cells

import "github.com/alem-platform/ap"

func HDigits(w int) {
	for i := 0; i < w; i++ {
		for j := 0; j < 7; j++ {
			ap.PutRune(' ')
		}
		PutNumber(i + 1)
	}
	ap.PutRune('\n')
}

func VDigits(stage int, number int) {
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
