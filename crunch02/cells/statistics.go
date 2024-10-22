package cells

func PrintStatistics(h, w, count, countM int) {
	PrintWords("Your Statistics:\n")
	PrintWords("Field Size: ")
	PrintWords(Num(h))
	PrintWords("x")
	PrintWords(Num(w))
	PrintWords("\nNumber of Bombs: ")
	PrintWords(Num(count))
	PrintWords("\nNumber of moves: ")
	PrintWords(Num(countM))
	PrintWords("\n")
}
