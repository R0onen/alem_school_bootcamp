package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	arg := os.Args[1:]
	for _, val := range arg {
		for _, v := range val {
			if !(v >= '0' && v <= '9') {
				return
			}
		}
	}
	if len(arg) > 1 || len(arg) == 0 {
		return
	}
	num := Atoi(string(arg[len(arg)-1]))
	Print(num)
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		Print(num)
	}
}

func Atoi(s string) int {
	c := 0

	for i, char := range s {
		if !(char >= '0' && char <= '9') {
			if !(i == 0 && char == '+' || char == '-') {
				return 0
			}
		} else {
			c = c*10 + int(char-'0')
		}
	}

	if s[0] == '-' {
		return -c
	}
	return c
}

func Print(n int) {
	s := Itoa(n)
	for _, v := range s {
		ap.PutRune(rune(v))
	}
	ap.PutRune('\n')
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	if n < 0 {
		s += "-"
		n = -n
	}
	var digits []byte
	for n != 0 {
		digits = append(digits, byte(n%10)+'0')
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		s += string(digits[i])
	}
	return s
}
