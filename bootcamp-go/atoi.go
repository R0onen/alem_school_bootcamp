package bootcamp

// import "fmt"

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	isNegative := false
	start := 0

	if s[0] == '-' {
		isNegative = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}


	result := 0

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

		result = result * 10 + int(s[i] - '0')
	}

	if isNegative {
		result = -result
	}

	return result


}

// func main() {
//     fmt.Println(Atoi("123"))
//     fmt.Println(Atoi("+123"))
//     fmt.Println(Atoi("-123"))
//     fmt.Println(Atoi("-123!"))
//     fmt.Println(Atoi("abc"))
// }