package bootcamp

func CountWords(s string) map[string]int {
	wordCount := make(map[string]int)
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			word += string(s[i] + 32)
		} else if s[i] >= 'a' && s[i] <= 'z' {
			word += string(s[i])
		} else if word != "" {
			wordCount[word]++
			word = ""
		}
	}
	if word != "" {
		wordCount[word]++
	}
	return wordCount
}
