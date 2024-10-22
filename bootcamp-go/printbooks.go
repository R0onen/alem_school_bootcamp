package bootcamp

import (
	"github.com/alem-platform/ap"
)

func PrintBooks(books []*Book) {
	LogN := func(n, base int) int {
		count := 0
		if n < 0 {
			n *= -1
			count += 1
		}
		if n < base {
			return 1
		}
		for n >= base {
			n /= base
			count++
		}
		return count + 1
	}
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var Itoa func(n int) string
	Itoa = func(n int) string {
		var s string = ""
		if n < 0 {
			s = "-"
			n *= -1
		}
		if n < 10 {
			return string(rune(n + 48))
		}
		s = s[:] + Itoa(n / 10)[:]
		return s[:] + string(rune(n%10+48))
	}
	size := [3]int{4, 6, 4}

	for _, book := range books {
		size[0] = Max(size[0], len(book.Name))
		size[1] = Max(size[1], len(book.Author))
		size[2] = Max(size[2], LogN(book.Year, 10))
	}

	for _, c := range "Name" {
		ap.PutRune(c)
	}
	for i := -1; i < size[0]-4; i++ {
		ap.PutRune(' ')
	}
	for _, c := range "Author" {
		ap.PutRune(c)
	}
	for i := -1; i < size[1]-6; i++ {
		ap.PutRune(' ')
	}
	for _, c := range "Year" {
		ap.PutRune(c)
	}
	ap.PutRune('\n')

	for _, book := range books {
		for _, c := range book.Name {
			ap.PutRune(c)
		}
		for i := -1; i < size[0]-len(book.Name); i++ {
			ap.PutRune(' ')
		}
		for _, c := range book.Author {
			ap.PutRune(c)
		}
		for i := -1; i < size[1]-len(book.Author); i++ {
			ap.PutRune(' ')
		}
		for _, c := range Itoa(book.Year) {
			ap.PutRune(c)
		}
		ap.PutRune('\n')
	}
}
