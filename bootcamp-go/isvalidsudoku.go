package bootcamp

// Function to check if Sudoku is valid
func ValidSudoku(n [9][9]int) bool {
	for i := 0; i < 9; i++ {
		rowCheck := make(map[int]bool)
		colCheck := make(map[int]bool)
		gridCheck := make(map[int]bool)

		for j := 0; j < 9; j++ {
			// Check row
			if n[i][j] != 0 {
				if rowCheck[n[i][j]] {
					return false
				}
				rowCheck[n[i][j]] = true
			}

			// Check column
			if n[j][i] != 0 {
				if colCheck[n[j][i]] {
					return false
				}
				colCheck[n[j][i]] = true
			}

			// Check 3x3 subgrid
			gridRow := 3*(i/3) + j/3
			gridCol := 3*(i%3) + j%3
			if n[gridRow][gridCol] != 0 {
				if gridCheck[n[gridRow][gridCol]] {
					return false
				}
				gridCheck[n[gridRow][gridCol]] = true
			}
		}
	}
	return true
}
