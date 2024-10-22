package bootcamp

func Capitalize(s string) string {
	result := []byte(s)

	for i := 0; i < len(result); i++ {
		if i == 0 || result[i-1] == ' ' {
			if result[i] >= 'a' && result[i] <= 'z' {
				result[i] -= 32
			}
		} else {
			if result[i] >= 'A' && result[i] <= 'Z' {
				result[i] += 32
			}
		}
	}

	return string(result)
}
