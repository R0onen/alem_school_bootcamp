package bootcamp

func SliceMakeN(n int) []int {
	slicemake := make([]int, n)
	for i := 0; i < n; i++ {
		slicemake[i] = i
	}
	return slicemake
}
