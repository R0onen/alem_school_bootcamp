package firststruct

func (u User) PasswordReliability() string {
	criteriaMet := 0
	var hasUpper, hasLower, hasDigit, hasSpecialChar bool
	for _, char := range u.password {
		if char >= 'a' && char <= 'z' {
			hasLower = true
		}
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
		if char <= 'a' && char >= 'z' && char <= 'A' && char >= 'Z' && char <= '0' && char >= '9' {
			hasSpecialChar = true
		}
	}
	if len(u.password) >= 8 {
		criteriaMet++
	}
	if hasDigit == true {
		criteriaMet++
	}
	if hasLower == true {
		criteriaMet++
	}
	if hasUpper == true {
		criteriaMet++
	}
	if hasSpecialChar == true {
		criteriaMet++
	}

	if criteriaMet >= 4 {
		return "strong"
	} else if criteriaMet >= 3 && criteriaMet < 5 {
		return "medium"
	} else if criteriaMet >= 1 && criteriaMet < 3 {
		return "easy"
	} else if u.password == "" {
		return "undefined"
	}
	return ""
}
