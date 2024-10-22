package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	var num int


	fmt.Scanf("%d", &num)

	if num == 0{return} 

	var max int = -99999999999999999
	var min int = 999999999999999

	for i := 0; i < num; i++ {
		var curr int
		fmt.Scanf("%d", &curr)
		if min > curr {
			min = curr
		} 
		 if max < curr {
			max = curr
		}
	}

	PutNumber(max)
	ap.PutRune(' ')
	PutNumber(min)
}

func PutNumber(n int) {
	s := make([]rune, 0)
	if n < 0 {
		ap.PutRune('-')
		n = n * (-1)
	}

	for n > 0 {

		mydigit := n % 10
		s = append(s, '0'+rune(mydigit))
		n = n / 10
	}

	for i := len(s) - 1; i >= 0; i-- {
		ap.PutRune(s[i])
	}
}
