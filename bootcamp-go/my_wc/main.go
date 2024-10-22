package main

import (
	//"bufio"
	"os"

	"github.com/alem-platform/ap"
)

func printf(s string) {
	for i := 0; i < len(s); i++ {
		ap.PutRune(rune(s[i]))
	}
}

func itoa(num int) string {
	if num == 0 {
		return "0"
	}

	var result []rune
	for num > 0 {
		remainder := num % 10
		result = append(result, rune(remainder+'0'))
		num /= 10
	}

	// cnahges the oreder of a string
	start := 0
	end := len(result) - 1
	for start < end {

		temp := result[start]
		result[start] = result[end]
		result[end] = temp

		start++
		end--
	}

	return string(result)
}

func main() {
	var args []string = os.Args
	var filename string
	inWord := false
	var linesCount, wordCount, byteCount int = 0, 0, 0

	filename = args[len(args)-1]
	data, err := os.ReadFile(filename)
	if err != nil {

		printf("my_wc: ")
		printf(filename)
		printf(": No such file or directory")

	}
	// if only file name was written
	if len(args) == 2 && err == nil {

		// count number of lines
		for i := 0; i < len(string(data)); i++ {

			//line count
			if data[i] == '\n' {
				linesCount++
			}

			// word count
			if data[i] != ' ' && data[i] != '\n' && data[i] != '\t' {
				if !inWord {
					wordCount++
					inWord = true
				}
			} else {
				inWord = false
			}
		}

		// byte count
		byteCount = len(data)

		// print result
		printf(itoa(linesCount))
		ap.PutRune(' ')
		printf(itoa(wordCount))
		ap.PutRune(' ')
		printf(itoa(byteCount))
		ap.PutRune(' ')
		printf(filename)

		// if no arguments were written
	} else if len(args) == 3 && err == nil {

		var options string
		options = args[1]

		for i := 1; i < len(options); i++ {

			if options[i] == 'l' {
				for i := 0; i < len(string(data)); i++ {
					if data[i] == '\n' {
						linesCount++
					}
				}
				printf(itoa(linesCount))
				ap.PutRune(' ')
			}

			if options[i] == 'w' {

				for i := 0; i < len(string(data)); i++ {
					if data[i] != ' ' && data[i] != '\n' && data[i] != '\t' {
						if !inWord {
							wordCount++
							inWord = true
						}
					} else {
						inWord = false
					}
				}

				printf(itoa(wordCount))
				ap.PutRune(' ')

			}

			if options[i] == 'c' {
				byteCount = len(data)

				printf(itoa(byteCount))
				ap.PutRune(' ')
			}
		}
		printf(filename)

	} else if len(args) == 1 {
		printf("usage: my_wc [-lwc] file")
	}

	ap.PutRune('\n')
}
