package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var text string

	reader := bufio.NewReader(os.Stdin)
	_ = reader
	_ = text

	fmt.Print("Enter text: ")
	text, _ = reader.ReadString('\n')

	text = Replace(text, "\n", "")

	countChr := getCountCharacters(text)
	fmt.Printf("Characters: %d\n", countChr)
	countSentences := getCountSentences(text)
	fmt.Printf("Sentences: %d\n", countSentences)

	countWords := getCountWords(text)
	fmt.Printf("Words: %d\n", countWords)

	countLetters := getCountLetters(text)
	fmt.Printf("Letters: %d\n", countLetters)

	countVowels := getCountVowels(text)
	fmt.Printf("Vowels: %d\n", countVowels)

	countConsonants := getCountConsonants(text)
	fmt.Printf("Consonants: %d\n", countConsonants)

	countDigits := getCountDigits(text)
	fmt.Printf("Digits: %d\n", countDigits)

	countSpaces := getCountSpaces(text)
	fmt.Printf("Spaces: %d\n", countSpaces)

	countSpecialChr := getCountSpecialCharacters(text)
	fmt.Printf("Special Characters: %d\n", countSpecialChr)
}

func Replace(s string, old string, new string) string {
	for i := 0; i+len(old) <= len(s); i++ {
		if s[i:i+len(old)] == old {
			s = s[:i] + new + s[i+len(old):]
			i += len(old)
		}
	}
	return s
}

// Check Symbol

func isLetter(l rune) bool {
	if 'a' <= l && l <= 'z' {
		return true
	}

	if 'A' <= l && l <= 'Z' {
		return true
	}

	return false
}

func isDigit(l rune) bool {
	if '0' <= l && l <= '9' {
		return true
	}

	return false
}

func isVowels(l rune) bool {
	for _, v := range "aeiou" {
		if v == l {
			return true
		}
	}

	for _, v := range "AEIOU" {
		if v == l {
			return true
		}
	}

	return false
}

// Get Data
func getCountSentences(text string) int {
	sentences := 0
	inSentence := false
	for i := 0; i < len(text); i++ {
		char := text[i]
		if char == '.' || char == '!' || char == '?' || char == '\n' {
			if inSentence {
				sentences++
				inSentence = false
			}
		} else if isLetter(rune(char)) {
			inSentence = true
		}
	}
	if inSentence {
		sentences++
	}

	return sentences
}

func getCountCharacters(text string) int {
	return len(text)
}

func getCountWords(s string) int {
	c := 0
	inWord := false
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanumeric := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for _, char := range s {
		if strings.ContainsRune(letters, char) {
			if !inWord {
				c++
				inWord = true
			}
		} else if !strings.ContainsRune(alphanumeric, char) {
			inWord = false
		}
	}
	return c
}

func getCountLetters(text string) int {
	count := 0

	for _, v := range text {
		if isLetter(v) {
			count++
		}
	}

	return count
}

func getCountVowels(text string) int {
	count := 0
	for _, v := range text {
		if isVowels(v) {
			count++
		}
	}

	return count
}

func getCountConsonants(text string) int {
	count := 0

	for _, v := range text {
		if !isVowels(v) && isLetter(v) {
			count++
		}
	}

	return count
}

func getCountDigits(text string) int {
	count := 0
	for _, v := range text {
		if isDigit(v) {
			count++
		}
	}
	return count
}

func getCountSpaces(text string) int {
	count := 0

	for _, v := range text {
		if v == ' ' {
			count++
		}
	}

	return count
}

func getCountSpecialCharacters(text string) int {
	specialCount := 0
	for _, char := range text {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && !unicode.IsSpace(char) {
			specialCount++
		}
	}
	return specialCount
}
