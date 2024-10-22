package bootcamp

// import "fmt"


func SliceSort(arr []int) {

	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] { 
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// func main() {
//     arr := []int{10, 90, 20, 5, 12, 3, 0}
//     SliceSort(arr)
//     fmt.Println(arr) // [0, 3, 5, 10, 12, 20, 90]
// }