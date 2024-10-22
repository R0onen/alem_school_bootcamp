package cells

import "github.com/alem-platform/ap"

// function of player cell
func Player(stage int, feature string) {
	ap.PutRune('|')
	if stage == 1 {
		for i := 0; i < 3; i++ {
			FillColor(Blue) // coloring rows
			ap.PutRune(' ')
		}
		ResetColor()
		ap.PutRune(rune(feature[1]))
		for i := 4; i < 7; i++ {
			FillColor(Blue)
			ap.PutRune(' ')
		}
		ResetColor()
	} else if stage == 2 {
		for i := 0; i < 7; i++ {
			FillColor(Blue)
			ap.PutRune('_')
		}
		ResetColor()
	} else {
		for i := 0; i < 7; i++ {
			FillColor(Blue)
			ap.PutRune(' ')
		}
		ResetColor()
	}
}
