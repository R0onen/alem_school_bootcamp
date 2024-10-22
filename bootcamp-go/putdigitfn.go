package bootcamp

import "github.com/alem-platform/ap"

func PutDigit(n int) {
	number := '0'
	if n >= 0 && n <= 9 {
		for i := 0; i < n; i++ {
			number++
		}
		ap.PutRune(number)
	}
}
