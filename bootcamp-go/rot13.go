package bootcamp

func Rot13(s string) string {
	result := []rune(s)

	for i := 0; i < len(result); i++ {
		if result[i] >= 'a' && result[i] <= 'z' {
			result[i] = 'a' + (result[i]-'a'+13)%26
		}

		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] = 'A' + (result[i]-'A'+13)%26
		}
	}

	return string(result)
}
