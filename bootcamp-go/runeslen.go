package bootcamp

func RunesLen(arr [20]rune) int {
	var count int
	for i := 0; i < 20; i++ {
		if arr[i] != '\x00' {
			count++
		} else {
			break
		}
	}
	return count
}
