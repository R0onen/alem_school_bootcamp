package bootcamp

import "github.com/alem-platform/ap"


func SliceMatrixPrint(matrix [][]rune) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			ap.PutRune(matrix[i][j])
			if j < len(matrix[i]) - 1 {
				ap.PutRune(' ')
			}
		}
		if i < len(matrix) - 1 {
			ap.PutRune('\n')
		}
	}
}

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
//   }
