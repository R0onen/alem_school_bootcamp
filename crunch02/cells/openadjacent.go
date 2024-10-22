package cells

func OpenAdjacentCells(h int, w int, field [][]rune, row int, col int, unrevealed_cells *int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			neighborH := row + i
			neighborW := col + j
			if neighborH >= 0 && neighborH < h && neighborW >= 0 && neighborW < w {
				// Only open if the cell is still unrevealed
				if field[neighborH][neighborW] == '.' {
					sum := CountAdjacentBombs(h, w, field, neighborH, neighborW)
					if sum > 0 {
						field[neighborH][neighborW] = rune('0' + sum)
					} else {
						field[neighborH][neighborW] = ' '
						// Recursively open adjacent cells
						OpenAdjacentCells(h, w, field, neighborH, neighborW, unrevealed_cells)
					}
				}
			}
		}
	}
	*unrevealed_cells--
}
