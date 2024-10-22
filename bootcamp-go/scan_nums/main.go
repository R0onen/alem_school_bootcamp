package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	sum := a + b
	subtraction := a - b
	multiplication := a * b
	division := a / b
	if b == 0 {
		ap.PutRune('F')
	}
	// sum
	if sum > 0 {
		digits := []rune{}
		for sum > 0 {
			digit := sum % 10
			digits = append([]rune{rune('0' + digit)}, digits...)
			sum /= 10
		}

		for _, digit := range digits {
			ap.PutRune(digit)
		}
		ap.PutRune(' ')
	} else if sum < 0 {
		sum = -sum
		ap.PutRune('-')
		digits := []rune{}
		for sum > 0 {
			digit := sum % 10
			digits = append([]rune{rune('0' + digit)}, digits...)
			sum /= 10
		}

		for _, digit := range digits {
			ap.PutRune(digit)
		}
		ap.PutRune(' ')
	} else {
		ap.PutRune('0')
		ap.PutRune(' ')
	}
	// subtraction
	if subtraction > 0 {
		digits := []rune{}
		for subtraction > 0 {
			digit := subtraction % 10
			digits = append([]rune{rune('0' + digit)}, digits...)
			subtraction /= 10
		}

		for _, digit := range digits {
			ap.PutRune(digit)
		}
		ap.PutRune(' ')
	} else if subtraction < 0 {
		subtraction = -subtraction
		ap.PutRune('-')
		digits := []rune{}
		for subtraction > 0 {
			digit := subtraction % 10
			digits = append([]rune{rune('0' + digit)}, digits...)
			subtraction /= 10
		}

		for _, digit := range digits {
			ap.PutRune(digit)
		}
		ap.PutRune(' ')
	} else {
		ap.PutRune('0')
		ap.PutRune(' ')
	}
	// multiplication
	if multiplication > 0 {
		digits := []rune{}
		for multiplication > 0 {
			digit := multiplication % 10
			digits = append([]rune{rune('0' + digit)}, digits...)
			multiplication /= 10
		}

		for _, digit := range digits {
			ap.PutRune(digit)
		}
		ap.PutRune(' ')
	} else if multiplication < 0 {
		multiplication = -multiplication
		ap.PutRune('-')
		digits := []rune{}
		for multiplication > 0 {
			digit := multiplication % 10
			digits = append([]rune{rune('0' + digit)}, digits...)
			multiplication /= 10
		}

		for _, digit := range digits {
			ap.PutRune(digit)
		}
		ap.PutRune(' ')
	} else {
		ap.PutRune('0')
		ap.PutRune(' ')
	}
	// division
	if b == 0 {
		ap.PutRune('F')
	} else {
		if division > 0 {
			digits := []rune{}
			for division > 0 {
				digit := division % 10
				digits = append([]rune{rune('0' + digit)}, digits...)
				division /= 10
			}

			for _, digit := range digits {
				ap.PutRune(digit)
			}
		} else if division < 0 {
			division = -division
			ap.PutRune('-')
			digits := []rune{}
			for division > 0 {
				digit := division % 10
				digits = append([]rune{rune('0' + digit)}, digits...)
				division /= 10
			}

			for _, digit := range digits {
				ap.PutRune(digit)
			}
		} else {
			ap.PutRune('0')
		}
	}
}
