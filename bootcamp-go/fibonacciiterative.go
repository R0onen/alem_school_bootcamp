package bootcamp

func FibonacciIterative(n int) int {
	terms := []int{0, 1}
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		terms = append(terms, terms[i-1]+terms[i-2])
	}
	return terms[n]
}
