package bootcamp

func MapContains(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}
