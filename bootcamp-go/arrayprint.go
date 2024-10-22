package bootcamp

import "github.com/alem-platform/ap"

func ArrayPrint(arr [20]int) {
	for i := 0; i < 20; i++ {
		digits := []rune{}
		isNegative := arr[i] < 0
		num := arr[i]

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
		if i != 19 {
			ap.PutRune(' ')
		}

	}
}
