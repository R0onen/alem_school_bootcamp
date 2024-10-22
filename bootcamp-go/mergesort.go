package bootcamp

func MergeSort(arr []int) {
	sorted(arr, 0, len(arr)-1)
}

func sorted(arr []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2
		sorted(arr, l, m)
		sorted(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m
	L := make([]int, n1)
	R := make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}
	i := 0
	j := 0
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
