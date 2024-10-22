package bootcamp 

// import "fmt"

func isUnique(s string) bool {
    
	seen := make(map[rune]bool)
	for _, char := range s {
        index := char - 'a' 
        if seen[index] {
            return false
        }
        seen[index] = true
    }
    return true
}


func UniqComb2(characters string) []string {

	if !isUnique(characters) || len(characters) < 2 {
		return []string{}
	}

	var combination []string

	for i := 0; i < len(characters); i++ {
		for j := 0; j < len(characters); j++ {
			if i != j {
				combination = append(combination, string(characters[i]) + string(characters[j]))
			}
		}
	}

	return combination
}

// func main() {
// 	fmt.Println(UniqComb2("abc"))  // ["ab", "ac", "ba", "bc", "ca", "cb"]
// 	fmt.Println(UniqComb2("ab"))   // ["ab", "ba"]
// 	fmt.Println(UniqComb2("a"))    // []
// 	fmt.Println(UniqComb2("aba"))  // []
// }


