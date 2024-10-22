package bootcamp

// import "fmt"

func IslandRemove(matrix [][]int, x, y int) {
	if x < 0 || y < 0 || x >= len(matrix[0]) || y >= len(matrix) || matrix[y][x] == 0 {
		return
	}
	matrix[y][x] = 0
	IslandRemove(matrix, x+1, y)
	IslandRemove(matrix, x, y+1)
	IslandRemove(matrix, x-1, y)
	IslandRemove(matrix, x, y-1)
}

// PrintMatrix prints the matrix
// func PrintMatrix(matrix [][]int) {
// 	for _, row := range matrix {
// 		for _, val := range row {
// 			fmt.Printf("%d ", val)
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	matrix := [][]int{
// 		{1, 1, 1, 0, 0, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 1, 0, 0, 0, 0},
// 		{0, 0, 0, 1, 2, 3, 1, 0, 0},
// 		{0, 0, 0, 1, 1, 1, 0, 0, 1},
// 		{0, 1, 0, 0, 0, 0, 0, 1, 2},
// 		{0, 0, 0, 1, 0, 0, 0, 0, 1},
// 		{0, 0, 1, 1, 2, 2, 0, 0, 0},
// 	}

// 	PrintMatrix(matrix)
// 	fmt.Println("Removing island at (4, 2):")
// 	IslandRemove(matrix, 4, 2)

// 	PrintMatrix(matrix)
// }
