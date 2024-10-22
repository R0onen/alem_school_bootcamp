package list

func (l *List) Len() int {
	if l.Head == nil {
		return 0
	}
	length := 1
	current := l.Head
	for current.Next != nil {
		length++
		current = current.Next
	}
	return length
}
