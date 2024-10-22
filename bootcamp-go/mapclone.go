package bootcamp

func MapClone(m map[string]int) map[string]int {
	newMap := make(map[string]int)
	if len(m) == 0 {
		return newMap
	} else {
		for key, value := range m {
			newMap[key] = value
		}
		return newMap
	}
}
