package bootcamp

func PowIterative(x int, power int) int {
	result := 1
	if power <= -1 {
		return -1
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return x
	}
	for i := 0; i < power; i++ {
		result *= x
	}
	return result
}
