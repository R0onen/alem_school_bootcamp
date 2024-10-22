package cells

import "github.com/alem-platform/ap"

func Cell() {
	ap.PutRune('|') // print left side
	for i := 0; i < 7; i++ {
		// FillColor(Red) // coloring rows
		ap.PutRune('X')
	}
	// ResetColor()
}
