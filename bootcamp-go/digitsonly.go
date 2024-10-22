package bootcamp

func DigitsOnly(s string) string {
	var digits string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits += string(s[i])
		}
	}
	return digits
}
