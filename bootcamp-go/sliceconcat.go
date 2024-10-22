package bootcamp

func SliceConcat(slices ...[]int) []int {
	var result []int
	for i := 0; i < len(slices); i++ {
		result = append(result, slices[i]...)
	}
	return result
}
