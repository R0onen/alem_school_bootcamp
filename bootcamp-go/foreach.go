package bootcamp

func ForEach(arr []int, fn func(*int)) {
	for i := 0; i < len(arr); i++ {
		fn(&arr[i])
	}
}
