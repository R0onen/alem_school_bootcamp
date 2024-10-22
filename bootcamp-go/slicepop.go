package bootcamp 

// import "fmt"

func SlicePop(arr *[]int) int {
	if len(*arr) == 0 {
		return 0
	} 
	
	lastElement := (*arr)[len(*arr) - 1]

	*arr = (*arr)[:len(*arr) -1]

	return lastElement
}

