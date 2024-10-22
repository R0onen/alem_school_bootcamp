package bootcamp

func MorseToText(s string) string {
	morseArr := make([]string, 0)
	var text string
	start := 0

	morseToText := map[string]string{
		".-":     "A",
		"-...":   "B",
		"-.-.":   "C",
		"-..":    "D",
		".":      "E",
		"..-.":   "F",
		"--.":    "G",
		"....":   "H",
		"..":     "I",
		".---":   "J",
		"-.-":    "K",
		".-..":   "L",
		"--":     "M",
		"-.":     "N",
		"---":    "O",
		".--.":   "P",
		"--.-":   "Q",
		".-.":    "R",
		"...":    "S",
		"-":      "T",
		"..-":    "U",
		"...-":   "V",
		".--":    "W",
		"-..-":   "X",
		"-.--":   "Y",
		"--..":   "Z",
		".----":  "1",
		"..---":  "2",
		"...--":  "3",
		"....-":  "4",
		".....":  "5",
		"-....":  "6",
		"--...":  "7",
		"---..":  "8",
		"----.":  "9",
		"-----":  "0",
		".-.-.-": ".",
		"--..--": ",",
		"..--..": "?",
	}

	s = s + " "
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			morseArr = append(morseArr, s[start:i])
			start = i + 1
		}
	}

	for i := 0; i < len(morseArr); i++ {
		text = text + morseToText[morseArr[i]]
	}

	return text
}
