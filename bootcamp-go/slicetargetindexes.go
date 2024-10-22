package bootcamp

func SliceTargetIndexes(arr []int, target int) []int {
	var targets []int
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			targets = append(targets, i)
		}
	}
	return targets
}
