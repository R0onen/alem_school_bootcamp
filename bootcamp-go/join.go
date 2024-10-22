package bootcamp

func Join(elements []string, sep string) string {
	var joinedstring string
	for i := 0; i < len(elements); i++ {
		joinedstring += elements[i]
		if i != len(elements)-1 {
			joinedstring += sep
		}
	}
	return joinedstring
}
