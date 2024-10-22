package cells

import (
	"github.com/alem-platform/ap"
)

// prints cell
func Cell(stage int) {
	ap.PutRune(rune('|'))
	if stage == 2 {
		for i := 0; i < 7; i++ {
			FillColor(White)
			ap.PutRune(rune('_'))
		}
		ResetColor()
	} else {
		for i := 0; i < 7; i++ {
			FillColor(White)
			ap.PutRune(rune(' '))
		}
		ResetColor()
	}
}
