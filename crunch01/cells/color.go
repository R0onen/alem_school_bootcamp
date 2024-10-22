package cells

import "github.com/alem-platform/ap"

func FillColor(color string) {
	for _, row := range color {
		ap.PutRune(rune(row))
	}
}

func ResetColor() {
	for _, row := range Reset {
		ap.PutRune(rune(row))
	}
}
