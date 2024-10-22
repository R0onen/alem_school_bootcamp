package list

func (l *List) Remove(n *ListNode) {
	if n == nil || l.Head == nil {
		return
	}
	if l.Head == n {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != n {
		current = current.Next
	}
	current.Next = current.Next.Next
}
