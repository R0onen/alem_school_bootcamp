package bootcamp

func ToLower(s string) string {
	var uppercase string
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			uppercase += string(int(s[i] + 32))
		} else {
			uppercase += string(s[i])
		}
	}
	return uppercase
}
