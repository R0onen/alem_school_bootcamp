package bootcamp

func MapMerge(m1 map[string]int, m2 map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range m1 {
		newMap[key] = value
	}
	for key, value := range m2 {
		newMap[key] = value
	}
	return newMap
}
