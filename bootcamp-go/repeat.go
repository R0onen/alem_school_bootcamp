package bootcamp

func Repeat(s string, count int) string {
	var repeatedstring string
	for i := 0; i < count; i++ {
		repeatedstring += s
	}
	return repeatedstring
}
