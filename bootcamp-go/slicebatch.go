package bootcamp


func SliceBatch(slice []int, size int) [][]int {

	if size <= 0 {
		return nil
	}

	var batches [][]int

	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice){
			end = len(slice)
		}
		
		batches = append(batches, slice[i:end])
	}

	return batches

}

