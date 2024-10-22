package bootcamp

func MapDeleteIf(m map[string]int, fn func(int) bool) {
	for key := range m {
		if fn(m[key]) {
			delete(m, key)
		}
	}
}
