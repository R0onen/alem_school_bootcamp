package bootcamp

import "github.com/alem-platform/ap"

func PrintBits(n byte) {
	for i := 7; i >= 0; i-- {
		if (n & (1 << i)) != 0 {
			ap.PutRune('1')
		} else {
			ap.PutRune('0')
		}
	}
	ap.PutRune('\n')
}
