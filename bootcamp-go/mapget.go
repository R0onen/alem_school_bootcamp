package bootcamp

func MapGet(m map[string]int, key string) int {
	value, ok := m[key]
	if ok == true {
		return value
	} else {
		return 0
	}
}
