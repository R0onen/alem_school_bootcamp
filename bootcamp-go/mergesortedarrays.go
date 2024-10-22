package bootcamp

func MergeSortedArrays(arr1, arr2 []int) []int {
	length := len(arr1) + len(arr2)
	arr3 := make([]int, length)
	j := len(arr1)
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			arr3[k] = arr1[i]
			i++
		} else {
			arr3[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		arr3[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		arr3[k] = arr2[j]
		j++
		k++
	}
	return arr3
}
