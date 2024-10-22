package bootcamp

// import "github.com/alem-platform/ap"

func SliceMatrixRotateClockwise(matrix [][]rune) {
	n := len(matrix)
	if n == 0 || n != len(matrix[0]) {
		return 
	}

	rotated := make([][]rune, n)
	for i := range rotated {
		rotated[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-i-1] = matrix[i][j]
		}
	}
	for i := range rotated {
		for j := range rotated[i] {
			matrix[i][j] = rotated[i][j]
		}
	}
}

// func SliceMatrixPrint(matrix [][]rune) {
// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[i]); j++ {
// 			ap.PutRune(matrix[i][j])
// 			if j < len(matrix[i]) - 1 {
// 				ap.PutRune(' ')
// 			}
// 		}
// 		if i < len(matrix) - 1 {
// 			ap.PutRune('\n')
// 		}
// 	}
// }

// func main() {
// 	var matrix = [][]rune{
// 	  { 'a', 'b', 'c' },
// 	  { 'd', 'e', 'f' },
// 	  { 'g', 'h', 'i' },
// 	}
  
// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// a b c
// 	// d e f
// 	// g h i
  
// 	SliceMatrixRotateClockwise(matrix)
  
// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// g d a
// 	// h e b
// 	// i f c
//   }