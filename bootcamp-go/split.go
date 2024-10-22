package bootcamp

func Split(s string, sep string) []string {
	var arr []string
	if sep == "" {
		for _, letter := range s {
			arr = append(arr, string(letter))
		}
	} else {
		start := 0
		for i := 0; i < len(s)-len(sep); i++ {
			if s[i:i+len(sep)] == sep {
				arr = append(arr, s[start:i])
				start = i + len(sep)
				i = start - 1
			}
		}

		if start < len(s) {
			arr = append(arr, s[start:])
		}
	}
	return arr
}
