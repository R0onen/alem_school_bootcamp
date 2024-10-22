package game_functions

func NextTick(current [][]rune, height int, width int, edgesflag bool, age [][]int) [][]rune {
	next := make([][]rune, height)
	for i := 0; i < height; i++ {
		next[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			liveNeighbors := CountLiveNeighbors(current, i, j, height, width, edgesflag)

			// Apply rules of the Game of Life
			if current[i][j] == '#' {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					next[i][j] = '.'
				} else {
					next[i][j] = '#'
				}
				age[i][j]++
			} else {
				if liveNeighbors == 3 {
					next[i][j] = '#'
				} else {
					next[i][j] = '.'
				}
				age[i][j] = 0
			}
		}
	}
	return next
}

func CountLiveNeighbors(grid [][]rune, x, y int, height int, width int, edgesPortal bool) int {
	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	count := 0
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if edgesPortal {
			nx = (nx + height) % height
			ny = (ny + width) % width
		}
		if nx >= 0 && nx < height && ny >= 0 && ny < width && grid[nx][ny] == '#' {
			count++
		}
	}
	return count
}

func NoLiveCells(g [][]rune) bool {
	for _, row := range g {
		for _, cell := range row {
			if cell == '#' {
				return false
			}
		}
	}
	return true
}

func GridsAreEqual(g1, g2 [][]rune, height int, width int) bool {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if g1[i][j] != g2[i][j] {
				return false
			}
		}
	}
	return true
}

func CountLiveCells(g [][]rune) int {
	count := 0
	for _, row := range g {
		for _, cell := range row {
			if cell == '#' {
				count++
			}
		}
	}
	return count
}
