package bootcamp

// import "fmt"

func NextPrime(n int) int {
	if n < 2 {
		return 2
	}

	for {
		n++
		if IsPrime(n) {
			return n
		}
	}
}

// func main() {
//     fmt.Println(NextPrime(10)) // 11
//     fmt.Println(NextPrime(11)) // 13
//     fmt.Println(NextPrime(-1)) // 2
// }
