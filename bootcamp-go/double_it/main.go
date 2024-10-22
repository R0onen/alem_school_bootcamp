package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	var L int
	fmt.Scanf("%d", &L)
	numbers := make([]int, L)
	for i := 0; i < L; i++ {
		fmt.Scanf("%d", &numbers[i])
	}
	for i := 0; i < L; i++ {
		numbers[i] *= 2
	}

	for i := 0; i < L; i++ {
		digits := []rune{}
		isNegative := numbers[i] < 0
		num := numbers[i]

		if isNegative {
			num = -num
		}

		if num == 0 {
			digits = append(digits, '0')
		} else {
			for num > 0 {
				digit := num % 10
				digits = append([]rune{rune('0' + digit)}, digits...)
				num /= 10
			}
		}

		if isNegative {
			ap.PutRune('-')
		}
		for _, digit := range digits {
			ap.PutRune(digit)
		}
		if i != L-1 {
			ap.PutRune(' ')
		}
	}
}
