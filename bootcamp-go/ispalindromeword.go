package bootcamp

// import "fmt"

func IsPalindromeWord(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') && !(s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}

	for i := 0; i < len(s)/2; i++ {
		left := s[i]
		right := s[len(s)-i-1]

		if left >= 'A' && left <= 'Z' {
			left += 'a' - 'A'
		}
		if right >= 'A' && right <= 'Z' {
			right += 'a' - 'A'
		}

		if left != right {
			return false
		}
	}

	return true
}

// func main() {
//     fmt.Println(IsPalindromeWord("racecar")) // true
//     fmt.Println(IsPalindromeWord("level"))   // true
//     fmt.Println(IsPalindromeWord("salem"))   // false
// }
