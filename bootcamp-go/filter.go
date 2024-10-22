package bootcamp

func Filter(arr []int, fn func(int) bool) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		if fn(arr[i]) == true {
			result = append(result, arr[i])
		}
	}
	return result
}
