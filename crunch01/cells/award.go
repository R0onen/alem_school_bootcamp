package cells

import "github.com/alem-platform/ap"

// function of award cell
func Award(stage int, feature string) {
	ap.PutRune('|')
	if stage == 1 {
		for i := 0; i < 3; i++ {
			FillColor(Yellow)
			ap.PutRune(' ')
		}
		ResetColor()
		ap.PutRune(rune(feature[2]))
		for i := 4; i < 7; i++ {
			FillColor(Yellow)
			ap.PutRune(' ')
		}
		ResetColor()
	} else if stage == 2 {
		for i := 0; i < 7; i++ {
			FillColor(Yellow)
			ap.PutRune('_')
		}
		ResetColor()
	} else {
		for i := 0; i < 7; i++ {
			FillColor(Yellow)
			ap.PutRune(' ')
		}
		ResetColor()
	}
}
