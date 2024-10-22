package bootcamp

func PowRecursion(x int, power int) int {
	if power <= -1 {
		return -1
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return x
	}
	return x * PowRecursion(x, power-1)
}
