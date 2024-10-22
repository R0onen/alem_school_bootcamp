package bootcamp

func Some(arr []int, fn func(int) bool) bool {
	ok := false
	for i := 0; i < len(arr); i++ {
		if fn(arr[i]) == true {
			ok = true
			break
		} else {
			ok = false
		}
	}
	return ok
}
