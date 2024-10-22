package cells

import "github.com/alem-platform/ap"

const (
	Reset  = "\033[0m"
	Red    = "\033[41m"
	White  = "\033[47m"
	Blue   = "\033[44m"
	Yellow = "\033[43m"
)

// rendering map with unrevealed cells
func Renderfirsttime(h int, w int, field [][]rune) {
	HDigits(w)
	Roof(w)
	for row := 0; row < h; row++ {
		for stage := 0; stage < 3; stage++ {
			VDigits(stage, row+1)
			for column := 0; column < w; column++ {
				switch field[row][column] {
				case '.':
					Cell()
				case '*':
					Cell()
				default:
					PrintWords("Error: Incorrect Input. Only . and *\n")
				}
			}
			ap.PutRune('|')
			ap.PutRune('\n')
		}
	}
}

// rendering map with opening every cell separately
func Render(h int, w int, field [][]rune, count *int) {
	HDigits(w)
	Roof(w)
	for row := 0; row < h; row++ {
		for stage := 0; stage < 3; stage++ {
			VDigits(stage, row+1)
			for column := 0; column < w; column++ {
				switch field[row][column] {
				case '.':
					Cell()
				case '*':
					Cell()
				default:
					if field[row][column] == ' ' {
						EmptyCell(stage)
					} else {
						CellData(stage, int(field[row][column]-'0'))
					}
				}
			}
			ap.PutRune('|')
			ap.PutRune('\n')
		}
	}
}

// rendering map with all opened bombs there
func Allbombs(h int, w int, field [][]rune, count *int) {
	HDigits(w)
	Roof(w)
	for row := 0; row < h; row++ {
		for stage := 0; stage < 3; stage++ {
			VDigits(stage, row+1)
			for column := 0; column < w; column++ {
				switch field[row][column] {
				case '.':
					CellData(stage, *count)
				case '*':
					Bomb(stage)
				default:
					if field[row][column] == ' ' {
						EmptyCell(stage)
					} else {
						CellData(stage, int(field[row][column]-'0'))
					}
				}
			}
			ap.PutRune('|')
			ap.PutRune('\n')
		}
	}
}
