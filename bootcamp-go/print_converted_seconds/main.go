package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func printNumber(num int) {
	if num == 0 {
		ap.PutRune('0')
		return
	}
	if num < 0 {
		ap.PutRune('-')
		num = -num
	}

	// Convert number to string manually
	var digits []rune
	for num > 0 {
		digits = append(digits, rune(num%10+'0'))
		num /= 10
	}

	// Print digits in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		ap.PutRune(digits[i])
	}
}

func main() {
	if len(os.Args) < 2 {
		ap.PutRune('N')
		ap.PutRune('V')
		ap.PutRune('\n')
		return
	}

	// Read the input argument as a string
	input := os.Args[1]

	// Convert input to int32 without using strconv
	seconds := 0
	for _, r := range input {
		if r < '0' || r > '9' {
			ap.PutRune('N')
			ap.PutRune('V')
			ap.PutRune('\n')
			return
		}
		seconds = seconds*10 + int(r-'0')
		if seconds < 0 { // Check for overflow
			ap.PutRune('N')
			ap.PutRune('V')
			ap.PutRune('\n')
			return
		}
	}

	// Limit check for int32 max value
	if seconds < 0 || seconds > 2147483647 {
		ap.PutRune('N')
		ap.PutRune('V')
		ap.PutRune('\n')
		return
	}

	// Convert seconds to days, hours, minutes, and seconds
	totalSeconds := seconds
	days := totalSeconds / 86400
	totalSeconds %= 86400
	hours := totalSeconds / 3600
	totalSeconds %= 3600
	minutes := totalSeconds / 60
	seconds = totalSeconds % 60

	// Print the output
	if days > 0 {
		printNumber(days)
		ap.PutRune('d')
		ap.PutRune(' ')
	}
	if hours > 0 {
		printNumber(hours)
		ap.PutRune('h')
		ap.PutRune(' ')
	}
	if minutes > 0 {
		printNumber(minutes)
		ap.PutRune('m')
		ap.PutRune(' ')
	}
	if seconds > 0 || (days == 0 && hours == 0 && minutes == 0) {
		printNumber(seconds)
		ap.PutRune('s')
	}
	ap.PutRune('\n')
}
