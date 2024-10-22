package bootcamp

func Trim(s string) string {
	start := 0
	end := len(s) - 1

	for start <= end && s[start] == ' ' {
		start++
	}

	for end >= start && s[end] == ' ' {
		end--
	}

	return s[start : end+1]
}
