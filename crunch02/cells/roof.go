package cells

import "github.com/alem-platform/ap"

func Roof(w int) {
	for i := 0; i < 4; i++ {
		ap.PutRune(' ')
	}
	for i := 0; i < 7*w+w-1; i++ {
		ap.PutRune('_')
	}
	ap.PutRune(' ')
	ap.PutRune('\n')
}
