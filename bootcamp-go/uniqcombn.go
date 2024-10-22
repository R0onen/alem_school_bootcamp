package bootcamp

// import "fmt"


// Helper function to generate combinations
func generateCombinations(current string, characters string, n int, combinations *[]string) {
	if len(current) == n {
		*combinations = append(*combinations, current)
		return
	}
	for i := 0; i < len(characters); i++ {
		generateCombinations(current+string(characters[i]), characters[:i]+characters[i+1:], n, combinations)
	}
}

// UniqCombN returns all unique n-character combinations
func UniqCombN(characters string, n int) []string {
	if !isUnique(characters) || len(characters) < n || n <= 0 {
		return []string{}
	}

	var combinations []string
	generateCombinations("", characters, n, &combinations)
	return combinations
}

// func main() {
// 	fmt.Println(UniqCombN("abc", 2))  // ["ab", "ac", "ba", "bc", "ca", "cb"]
// 	fmt.Println(UniqCombN("abcd", 3)) // ["abc", "abd", "acb", "acd", "adb", "adc", "bac", "bad", "bca", "bcd", "cab", "cad", "dab", "dac"]
// 	fmt.Println(UniqCombN("ab", 2))   // ["ab", "ba"]
// 	fmt.Println(UniqCombN("a", 1))    // ["a"]
// 	fmt.Println(UniqCombN("aba", 2))  // []
// 	fmt.Println(UniqCombN("abc", 4))  // []
// }
