package bootcamp

func IsSorted(arr []int, fn func(int, int) bool) bool {
	for i := 0; i < len(arr)-1; i++ {
		if !fn(arr[i], arr[i+1]) {
			return false
		}
	}
	return true
}
