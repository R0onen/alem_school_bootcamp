package game_functions

func IsValidRow(s string) bool {
	for _, r := range s {
		if r != '#' && r != '.' {
			return false
		}
	}
	return true
}

func IsNumber(s string) bool {
	for l := 0; l < len(s); l++ {
		if s[l] < '0' || s[l] > '9' {
			return false
		}
	}
	return true
}

func ToInt(s string) int {
	result := 0
	for _, r := range s {
		result = result*10 + int(r-'0')
	}
	return result
}
