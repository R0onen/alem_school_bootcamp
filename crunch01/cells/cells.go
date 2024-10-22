package cells

import (
	//"crunch01/cells"
	//"fmt"

	"github.com/alem-platform/ap"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[41m"
	White  = "\033[47m"
	Blue   = "\033[44m"
	Yellow = "\033[43m"
)

func Render(height int, width int, field [][]rune, feature string) {
	Letters(width)
	Roof(width)
	for row := 0; row < height; row++ {
		for stage := 0; stage < 3; stage++ {
			Digits(stage, row+1)
			for column := 0; column < width; column++ {
				switch field[row][column] {
				case '0':
					Wall(feature)
				case '1':
					Cell(stage)
				case '2':
					Player(stage, feature)
				case '3':
					Award(stage, feature)
				}
			}
			ap.PutRune('|')
			ap.PutRune('\n')
		}
	}
}
