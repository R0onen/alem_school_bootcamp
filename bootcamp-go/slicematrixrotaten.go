package bootcamp

// import "github.com/alem-platform/ap"


func SliceMatrixRotateN(matrix [][]rune, turns int) {
	n := len(matrix)
	if n == 0 || n != len(matrix[0]) {
		return 
	}


	turns = (turns % 4 + 4) % 4

	rotated := make([][]rune, n)
	for i := range rotated {
		rotated[i] = make([]rune, n)
	}

	for k := 0; k < turns; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				rotated[j][n-i-1] = matrix[i][j]
			}
		}

		for i := range matrix {
			copy(matrix[i], rotated[i])
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
  
// 	SliceMatrixRotateN(matrix, 1)
  
// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// g d a
// 	// h e b
// 	// i f c
  
// 	SliceMatrixRotateN(matrix, -2)
  
// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// c f i
// 	// b e h
// 	// a d g
//   }