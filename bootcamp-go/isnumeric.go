package bootcamp

func IsNumeric(s string) bool {
	ok := false
	if s[0] == '+' || s[0] == '-' {
		for i := 1; i < len(s); i++ {
			if s[i] != ' ' {
				if s[i] >= '0' && s[i] <= '9' {
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
	} else {
		for i := 0; i < len(s); i++ {
			if s[i] != ' ' {
				if s[i] >= '0' && s[i] <= '9' {
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
	}
	return ok
}
