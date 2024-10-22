package bootcamp

func Concat(s ...string) string {
	var singlestring string
	for _, elements := range s {
		singlestring += elements
	}
	return singlestring
}
