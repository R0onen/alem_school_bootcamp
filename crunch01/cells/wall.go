package cells

import "github.com/alem-platform/ap"

// print cell of wall
func Wall(feature string) {
	ap.PutRune('|') // print left side
	for i := 0; i < 7; i++ {
		FillColor(Red) // coloring rows
		ap.PutRune(rune(feature[0]))
	}
	ResetColor()
}
