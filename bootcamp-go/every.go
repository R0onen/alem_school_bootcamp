package bootcamp

func Every(arr []int, fn func(int) bool) bool {
	ok := false
	for i := 0; i < len(arr); i++ {
		if fn(arr[i]) == true {
			ok = true
		} else {
			ok = false
			break
		}
	}
	return ok
}
