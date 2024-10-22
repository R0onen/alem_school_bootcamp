package bootcamp

import (
	"bufio"
	"os"
)

type Book struct {
	Name   string
	Author string
	Year   int
}

func _Atoi(s string) int {
	IsNumeric := func(s string) bool {
		signed := false
		for _, c := range s {
			if !(c >= '0' && c <= '9') {
				if signed {
					return false
				}
				if c == '+' || c == '-' {
					signed = true
					continue
				}
				return false
			}
		}
		return true
	}
	var num int = 0
	if !IsNumeric(s) {
		return num
	}
	var negative bool = s[0] == '-'
	_ = negative
	for i, c := range s {
		if c == '-' || c == '+' {
			if i != 0 {
				return 0
			}
			continue
		}
		num = num*10 + int(c-'0')
		if negative {
			num -= int(c-'0') * 2
		}
	}
	return num
}

func GetBooksFromCsv(path string) []*Book {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ids := []int{0, 1, 2} // Name Autor Year
	var books []*Book
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		bookFields := make([]string, 3)
		fieldIndex := 0
		startIndex := 0

		for i := 0; i < len(line); i++ {
			if line[i] == ',' {
				bookFields[fieldIndex] = line[startIndex:i]
				startIndex = i + 1
				fieldIndex++
			}
		}
		bookFields[fieldIndex] = line[startIndex:]
		if i == 0 {
			for i, val := range bookFields {
				if val == "Name" {
					ids[0] = i
				}
				if val == "Author" {
					ids[1] = i
				}
				if val == "Year" {
					ids[2] = i
				}
			}
			i++
			continue
		}
		book := &Book{
			Name:   bookFields[ids[0]],
			Author: bookFields[ids[1]],
			Year:   _Atoi(bookFields[ids[2]]),
		}

		books = append(books, book)
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return books
}
