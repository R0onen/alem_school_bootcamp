package cells

func IsAllMines(field [][]rune) bool {
	for _, row := range field {
		for _, cell := range row {
			if cell != '*' {
				return false
			}
		}
	}
	return true
}
