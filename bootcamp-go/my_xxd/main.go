package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func printLine(s string) {
	for i := 0; i < len(s); i++ {
		ap.PutRune(rune(s[i]))
	}
}

func Atoi(s string) int {
	flag := false
	var m int
	var i int
	if s[0] == '-' {
		flag = true
		i = 1
	}
	if s[0] == '+' {
		i = 1
	}
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			m += int(s[i] - '0')
			if i < len(s)-1 {
				m *= 10
			}
		} else {
			return 0
		}
	}
	if flag {
		m *= -1
	}
	return m
}

func to16(number int) string {
	base := "0123456789abcdef"
	ans := ""
	for number != 0 {
		ans += string(base[number%16])
		number /= 16
	}
	if len(ans) < 2 {
		ans += "0"
	}
	reverse := ""
	for i := len(ans) - 1; i >= 0; i-- {
		reverse += string(ans[i])
	}
	return reverse
}

func intToHex(n int) string {
	hexDigits := "0123456789abcdef"
	if n == 0 {
		return "0"
	}

	hexStr := ""
	for n > 0 {
		hexStr = string(hexDigits[n%16]) + hexStr
		n /= 16
	}
	return hexStr
}

func incrementByHex(q int) string {
	hexValue := intToHex(q)
	var ans string
	for i := 0; i < 8-len(hexValue); i++ {
		ans += "0"
	}
	ans += hexValue
	return ans
}

func Separator(s string, n int) []string {
	var s_arr []string
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == ' ' {
	// 		s = s[:i] + s[i+1:]
	// 		i--
	// 	}
	// }
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			s = s[:i] + "." + s[i+1:]
		}
	}
	var Start int = 0

	for i := 0; i < len(s); i++ {
		if (i+1)%n == 0 {
			s_arr = append(s_arr, s[Start:i+1])
			Start = i + 1
		}
	}
	s_arr = append(s_arr, s[Start:])
	return s_arr
}

func solve(content []byte, size int) {
	j := 0
	sep := Separator(string(content), size)
	for i := 0; i < len(sep); i++ {
		var temp string
		if sep[i] == "" {
			continue
		}
		temp += incrementByHex(j) + ": "
		for ; j < len(content); j++ {
			if j == size*(i+1) {
				break
			}
			if j%2 == 0 && j != size*i && j != (size*i)-1 {
				temp += " "
			}
			temp += to16(int(content[j]))

		}
		if i == len(sep)-1 && len(content)%size != 0 {
			n := (size-(len(content)%size))*2 + (size-(len(content)%size))/2 // 28+7 = 35
			for k := 0; k < n; k++ {
				temp += " "
			}
		}
		temp += "  " + sep[i]
		printLine(temp)
		ap.PutRune('\n')
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		printLine("usage: my_xxd [-c number] file\n")
		return
	}

	c := false
	size := 16
	for i := 1; i < len(args)-1; i++ {
		if args[i] == "-c" {
			c = true
		}
		if c {
			size = Atoi(args[i])
		}
	}
	content, err := os.ReadFile(args[len(args)-1])
	if err != nil {
		printLine("my_xxd: " + args[len(args)-1] + ": No such file or directory\n")
		return
	}
	solve(content, size)
}
