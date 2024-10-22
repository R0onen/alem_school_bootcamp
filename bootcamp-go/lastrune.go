package bootcamp

func LastRune(s string) rune {
	var lastrune rune
	if s != "" {
		lastrune = rune(s[len(s)-1])
	} else {
		lastrune = rune(0)
	}
	return lastrune
}
