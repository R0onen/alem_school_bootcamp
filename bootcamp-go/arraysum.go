package bootcamp

func ArraySum(arr [20]int) int {
	sum := 0
	for i := 0; i < 20; i++ {
		sum += arr[i]
	}
	return sum
}
