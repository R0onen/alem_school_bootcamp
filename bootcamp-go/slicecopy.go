package bootcamp

func SliceCopy(dst, src []int) {
	if len(dst) > len(src) {
		for i := 0; i < len(src); i++ {
			dst[i] = src[i]
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = src[i]
		}
	}
}
