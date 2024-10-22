package bootcamp

func ToUpper(s string) string {
	var uppercase string
	for i := 0; i < len(s); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			uppercase += string(int(s[i] - 32))
		} else {
			uppercase += string(s[i])
		}
	}
	return uppercase
}
