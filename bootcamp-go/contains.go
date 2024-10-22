package bootcamp

func Contains(str string, substr string) bool {
	if substr == "" {
		return true
	}
	return index(str, substr) != -1
}

func index(str, substr string) int {
	lenStr := len(str)
	lenSub := len(substr)

	for i := 0; i <= lenStr-lenSub; i++ {
		if str[i:i+lenSub] == substr {
			return i
		}
	}
	return -1
}
