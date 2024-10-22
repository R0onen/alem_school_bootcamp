package list

func (l *List) Reverse() {
	current := l.Head
	var prev *ListNode = nil
	var next *ListNode = nil

	l.Tail, l.Head = l.Head, l.Tail
	for current != nil {
		next = current.Next
		current.Next = prev

		prev = current
		current = next
	}

	l.Head = prev
}
