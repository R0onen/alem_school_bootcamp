package bootcamp

// import "fmt"

func SliceInsert(arr *[]int, idx int, values ...int) bool {
	if idx < 0 || idx > len(*arr) {
		return false
	}

	ls := make([]int, 0, len(*arr)+len(values))

	ls = append(ls, (*arr)[:idx]...)

	ls = append(ls, values...)

	ls = append(ls, (*arr)[idx:]...)

	*arr = ls
	
	return true
}

// func main() {
//     arr := []int{1}
//     fmt.Println(arr) // [1]

//     fmt.Println(SliceInsert(&arr, 0, 10))
//     fmt.Println(arr) // [10, 1]

//     fmt.Println(SliceInsert(&arr, len(arr), 20))
//     fmt.Println(arr) // [10, 1, 20]

//     fmt.Println(SliceInsert(&arr, 2, 3)) // true
//     fmt.Println(arr) // [10, 1, 3, 20]

//     fmt.Println(SliceInsert(&arr, -1, 100)) // false
//     fmt.Println(arr) // [10, 1, 3, 20]
// }