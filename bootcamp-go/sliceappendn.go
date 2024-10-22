package bootcamp

func SliceAppendN(n int) []int {
	var sliceappend []int
	for i := 0; i < n; i++ {
		sliceappend = append(sliceappend, i)
	}
	return sliceappend
}
