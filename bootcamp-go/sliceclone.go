package bootcamp

func SliceClone(src, dst *[]int) {
	leng := len(*src)
	capa := cap(*src)

	*dst = make([]int, leng, capa)

	for i := 0; i < leng; i++ {
		(*dst)[i] = (*src)[i]
	}
}
