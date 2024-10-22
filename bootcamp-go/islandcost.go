package bootcamp 

// import "fmt"

func IslandCost(matrix [][]int, x, y int) int {
	if x < 0 || y < 0 || y >= len(matrix) || x >= len(matrix[0]) || matrix[y][x] == 0 {
		return 0
	}

	counter := matrix[y][x]
	matrix[y][x] = 0

	return counter + IslandCost(matrix, x-1, y) + IslandCost(matrix, x+1, y) + IslandCost(matrix, x, y-1) + IslandCost(matrix, x, y+1)
}

// func main() {
// 	matrix := [][]int{
// 	  { 1, 1, 1, 0, 0, 0, 0, 0, 0 },
// 	  { 1, 1, 0, 0, 1, 0, 0, 0, 0 },
// 	  { 0, 0, 0, 1, 2, 3, 1, 0, 0 },
// 	  { 0, 0, 0, 1, 1, 1, 0, 0, 1 },
// 	  { 0, 1, 0, 0, 0, 0, 0, 1, 2 },
// 	  { 0, 0, 0, 1, 0, 0, 0, 0, 1 },
// 	  { 0, 0, 1, 1, 2, 2, 0, 0, 0 },
// 	}
  
// 	fmt.Println(IslandCost(matrix, 4, 2)) // 11
// 	fmt.Println(IslandCost(matrix, 0, 0)) // 5
// 	fmt.Println(IslandCost(matrix, 1, 4)) // 1
// 	fmt.Println(IslandCost(matrix, 0, 3)) // 0
//   }