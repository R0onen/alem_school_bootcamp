package bootcamp

func RomanToInt(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		currentstr := romans[string(s[i])]
		if i+1 < len(s) && currentstr < romans[string(s[i+1])] {
			sum -= currentstr
		} else {
			sum += currentstr
		}
	}
	return sum
}
