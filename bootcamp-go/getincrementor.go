package bootcamp

func GetIncrementor(start int, step int) func() int {
	current := start
	return func() int {
		current += step
		return current
	}
}
