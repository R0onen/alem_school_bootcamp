package bootcamp

func EightQueens() [][]int {
	var solutions [][]int
	columns := make([]int, 8)
	placeQueens(0, columns, &solutions)

	// Sort solutions using custom MergeSort implementation
	MergeSortSolutions(&solutions)

	for i := range solutions {
		for j := range solutions[i] {
			solutions[i][j]++
		}
	}

	return solutions
}

func placeQueens(row int, columns []int, solutions *[][]int) {
	if row == 8 {
		// Found a solution, add to solutions list
		solution := make([]int, 8)
		copy(solution, columns)
		*solutions = append(*solutions, solution)
		return
	}
	for col := 0; col < 8; col++ {
		columns[row] = col
		if isValid(columns, row) {
			placeQueens(row+1, columns, solutions)
		}
	}
}

func isValid(columns []int, row int) bool {
	for r := 0; r < row; r++ {
		diff := columns[r] - columns[row]
		if diff == 0 || diff == row-r || diff == r-row {
			return false
		}
	}
	return true
}

func MergeSortSolutions(arr *[][]int) {
	aux := make([][]int, len(*arr))
	copy(aux, *arr)
	mergeSortHelper(aux, *arr, 0, len(*arr))
}

func mergeSortHelper(src, dest [][]int, low, high int) {
	if high-low < 2 {
		return
	}
	mid := (low + high) / 2
	mergeSortHelper(dest, src, low, mid)
	mergeSortHelper(dest, src, mid, high)
	Merge(src, dest, low, mid, high)
}

func Merge(src, dest [][]int, low, mid, high int) {
	i, j := low, mid

	for k := low; k < high; k++ {
		if i < mid && (j >= high || compareSolutions(src[i], src[j])) {
			dest[k] = src[i]
			i++
		} else {
			dest[k] = src[j]
			j++
		}
	}
}

func compareSolutions(sol1, sol2 []int) bool {
	for i := 0; i < 8; i++ {
		if sol1[i] < sol2[i] {
			return true
		} else if sol1[i] > sol2[i] {
			return false
		}
	}
	return false
}
