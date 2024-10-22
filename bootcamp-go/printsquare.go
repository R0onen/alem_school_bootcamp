package bootcamp

import "github.com/alem-platform/ap"

func PrintSquare(w int) {
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			ap.PutRune('0')
			if j != w-1 {
				ap.PutRune(' ')
			}
		}
		ap.PutRune('\n')
	}
}
