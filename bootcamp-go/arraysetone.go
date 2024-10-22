package bootcamp

func ArraySetOne(arr *[20]int, idx int, val int) bool {
	ok := false
	if idx < 0 || idx > 20 {
		return ok
	} else {
		arr[idx] = val
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == val && i == idx {
			ok = true
			break
		}
	}
	return ok
}
