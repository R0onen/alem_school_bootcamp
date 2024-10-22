package list

func (l *List) Back() *ListNode {
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current
}
