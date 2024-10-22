package bootcamp

// import "fmt"

func RoundBrackets(s string) bool {
	count := 0

	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '(' {
			count++
		}

		if char == ')' {
			count--

			if count < 0 {
				return false
			}
		}

	}

	return count == 0
}

// func main() {
// 	fmt.Println(RoundBrackets("()()()"))   // true
// 	fmt.Println(RoundBrackets("(salem)"))  // true
// 	fmt.Println(RoundBrackets(")(1)(f)(")) // false
// 	fmt.Println(RoundBrackets("))(()"))    // false
// 	fmt.Println(RoundBrackets(""))         // false
// }
