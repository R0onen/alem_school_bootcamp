package bootcamp

func SliceGetClone(src []int) []int {
	leng := len(src)
	capa := cap(src)

	dst := make([]int, leng, capa)

	for i := 0; i < leng; i++ {
		dst[i] = src[i]
	}
	return dst
}
