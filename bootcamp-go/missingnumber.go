package bootcamp

func MissingNumber(arr []int) int {
	aldik := make([]bool, len(arr)+1)
	for _, i := range arr {
		if i > 0 && i <= len(arr) {
			aldik[i] = true
		}
	}
	for i := 1; i < len(aldik); i++ {
		if !aldik[i] {
			return i
		}
	}
	return len(arr) + 1
}
