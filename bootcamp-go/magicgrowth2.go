package bootcamp 

// import "fmt"

func MagicGrowth2() []string {
	var result []string

	for i := 0; i <= 8; i++ {
		for j := i + 1; j <= 9; j++ {
			result = append(result, Itoa(i) + Itoa(j))
		}
	}

	return result
}

// func Itoa(num int) string {
// 	if num == 0 {
// 		return "0"
// 	}

// 	isNegative := num < 0 // True or False
// 	if isNegative {
// 		num = -num
// 	}

// 	var result []rune

// 	for num > 0 {
// 		remainder := num % 10
// 		result = append(result, rune(remainder + '0'))
// 		num /= 10
// 	}

// 	if isNegative {
// 		result = append(result, '-')
// 	}

	
// 	start := 0
// 	end := len(result) - 1
// 	for start < end {

// 		temp := result[start]
// 		result[start] = result[end]
// 		result[end] = temp

// 		start++
// 		end--
// 	}

// 	return string(result)
// }

// func main() {
// 	fmt.Println(MagicGrowth2())  // ["01", "02", ... "08", "09", "12", "13" ... "78", "79", "89"]
// }