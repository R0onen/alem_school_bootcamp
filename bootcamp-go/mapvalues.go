package bootcamp

func MapValues(m map[string]int) []int {
	var values []int
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
