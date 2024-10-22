package bootcamp

func IsAlpha(s string) bool {
	ok := false
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
				ok = true
			} else {
				ok = false
				break
			}
		} else {
			ok = false
			break
		}
	}
	return ok
}
