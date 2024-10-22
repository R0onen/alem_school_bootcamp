package bootcamp

import "github.com/alem-platform/ap"

func PutNumber(n int) {
	if n == 0 {
		ap.PutRune('0')
	}
	if n < 0 {
		ap.PutRune('-')
		n = -n
	}
	digits := []rune{}
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}

	for _, digit := range digits {
		ap.PutRune(digit)
	}
}
