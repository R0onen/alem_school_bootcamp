package bootcamp

func RotN(s string, n int) string {
	result := []rune(s)

	for i := 0; i < len(result); i++ {
		if result[i] >= 'a' && result[i] <= 'z' {
			result[i] = 'a' + rune((int(result[i])-int('a')+n)%26)
		}

		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] = 'A' + rune((int(result[i])-int('A')+n)%26)
		}
	}

	return string(result)
}
