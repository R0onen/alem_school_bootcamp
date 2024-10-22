package bootcamp

func MapKeys(m map[string]int) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
