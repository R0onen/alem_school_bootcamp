package bootcamp

func ArraySlice(arr [20]int, low int, high int) []int {
	if low >= 0 && high <= 20 && low < high {
		return arr[low:high]
	} else {
		return nil
	}
}
