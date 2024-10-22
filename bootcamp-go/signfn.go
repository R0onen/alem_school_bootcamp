package bootcamp

import "github.com/alem-platform/ap"

func Sign(n int) {
	if n > 0 {
		ap.PutRune('+')
	}
	if n < 0 {
		ap.PutRune('-')
	}
	if n == 0 {
		ap.PutRune('0')
	}
}
