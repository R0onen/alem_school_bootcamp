package bootcamp

func Find(arr []int, fn func(int) bool) *int {
	for i := 0; i < len(arr); i++ {
		if fn(arr[i]) {
			return &arr[i]
		}
	}
	return nil
}
