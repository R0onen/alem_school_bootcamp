package bootcamp

// import "fmt"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	counter := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			counter++
		}
	}

	if counter == 1 {
		return true
	} else {
		return false
	}
}

// func main() {
//     fmt.Println(IsPrime(11))  // true
//     fmt.Println(IsPrime(4))   // false
//     fmt.Println(IsPrime(-11)) // false
// }
