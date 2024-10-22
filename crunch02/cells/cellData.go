package cells

import "github.com/alem-platform/ap"

func CellData(stage int, count int) {
	color := GetColor(count)
	ap.PutRune('|')
	if stage == 1 {
		for i := 0; i < 3; i++ {
			ap.PutRune(' ')
		}
		FillColor(color)
		PrintWords(Num(count))
		ResetColor()
		for i := 4; i < 7; i++ {
			ap.PutRune(' ')
		}
	} else if stage == 2 {
		for i := 0; i < 7; i++ {
			ap.PutRune('_')
		}
	} else {
		for i := 0; i < 7; i++ {
			ap.PutRune(' ')
		}
	}
}
