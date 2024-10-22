package cells

func CountAdjacentBombs(h int, w int, field [][]rune, row int, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			neighborH := row + i
			neighborW := col + j
			if neighborH >= 0 && neighborH < h && neighborW >= 0 && neighborW < w {
				if field[neighborH][neighborW] == '*' {
					count++
				}
			}
		}
	}
	return count
}
