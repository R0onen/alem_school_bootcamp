package bootcamp

// import "fmt"

func ConvertNbrBase(n int, base string) string {
	if len(base) < 2 || hasDuplicateChars(base) || n == 0 {
		return ""
	}

	baseLength := len(base)
	var result []byte

	for n > 0 {
		remainder := n % baseLength
		result = append(result, base[remainder])
		n /= baseLength
	}

	reversed := make([]byte, len(result))
	for i := 0; i < len(result); i++ {
		reversed[i] = result[len(result)-1-i]
	}

	return string(reversed)
}

func hasDuplicateChars(base string) bool {
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return true
			}
		}
	}
	return false
}

// func main() {
// 	result := ConvertNbrBase(1465, "0123456789")
// 	fmt.Println(result) // 1465

// 	result = ConvertNbrBase(1465, "01")
// 	fmt.Println(result) // 10110111001

// 	result = ConvertNbrBase(1465, "01234567")
// 	fmt.Println(result) // 2671

// 	result = ConvertNbrBase(1465, "0123456789ABCDEF")
// 	fmt.Println(result) // 5B9

// 	result = ConvertNbrBase(1465, "00")
// 	fmt.Println(result) //
//   }
