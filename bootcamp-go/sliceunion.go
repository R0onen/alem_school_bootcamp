package bootcamp

// import "fmt"

func SliceUnion(slices ...[]int) []int {
	result := []int{}

	for i := 0; i < len(slices); i++ {
		slice := slices[i]

		for j := 0; j < len(slice); j++ {
			value := slice[j]
			if !contains(result, value){
				result = append(result, value)
			}
		}
	}

	return result
}


func contains(slice []int, value int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return true
		}
	}

	return false
}


// func main() {
//     result := SliceUnion([]int{1, 2, 3, 20}, []int{1, 2, 3, 4}, []int{15, 0, 2})
//     fmt.Println(result) // [1, 2, 3, 20, 4, 5, 15, 0]
// }