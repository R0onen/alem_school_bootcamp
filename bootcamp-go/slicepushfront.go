package bootcamp

func SlicePushFront(arr *[]int, values ...int) {
	result := make([]int, len(values)+len(*arr))
	for i := 0; i < len(values); i++ {
		result[i] = values[i]
	}
	for i := 0; i < len(*arr); i++ {
		result[len(values)+i] = (*arr)[i]
	}
	*arr = result
}
