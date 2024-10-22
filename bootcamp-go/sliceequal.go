package bootcamp

func SliceEqual(arr1, arr2 []int) bool {
	ok := false
	if len(arr1) != len(arr2) {
		return false
	} else {
		for i := 0; i < len(arr1); i++ {
			if arr1[i] == arr2[i] {
				ok = true
				continue
			} else {
				ok = false
				break
			}
		}
		return ok
	}
}
