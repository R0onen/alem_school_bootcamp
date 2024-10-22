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

func GetColor(count int) string {
	switch {
	case count == 1:
		return White
	case count == 2:
		return Yellow
	case count == 3:
		return Blue
	case count == 4:
		return Red
	default:
		return White
	}
}
