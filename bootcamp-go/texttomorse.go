package bootcamp

func TextToMorse(s string) string {
	var morse string
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			s = toUpper(s)
		}
	}
	morseCode := map[string]string{
		"A": ".-", "B": "-...", "C": "-.-.", "D": "-..",
		"E": ".", "F": "..-.", "G": "--.", "H": "....",
		"I": "..", "J": ".---", "K": "-.-", "L": ".-..",
		"M": "--", "N": "-.", "O": "---", "P": ".--.",
		"Q": "--.-", "R": ".-.", "S": "...", "T": "-",
		"U": "..-", "V": "...-", "W": ".--", "X": "-..-",
		"Y": "-.--", "Z": "--..",
		"1": ".----", "2": "..---", "3": "...--", "4": "....-",
		"5": ".....", "6": "-....", "7": "--...", "8": "---..",
		"9": "----.", "0": "-----",
		".": ".-.-.-", ",": "--..--", "?": "..--..",
	}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			morse += morseCode[string(s[i])]
			if i != len(s)-1 {
				morse += " "
			}
		}
	}
	return morse
}

func toUpper(s string) string {
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
