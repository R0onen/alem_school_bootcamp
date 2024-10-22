package bootcamp

func ConvertBaseNbr(n string, base string) int {
	if len(base) < 2 {
		return -1
	}
	charIndex := make(map[rune]int)
	for i, char := range base {
		charIndex[char] = i
	}
	result := 0
	baseLen := len(base)
	for _, digit := range n {
		if index, exists := charIndex[digit]; exists {
			result = result*baseLen + index
		} else {
			return -1
		}
	}
	return result
}
