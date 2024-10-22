package list

func (l *List) RemoveIf(fn func(n *ListNode) bool) {
	node := l.Head
	for node != nil {
		if fn(node) {
			l.Remove2(node)
		}
		node = node.Next
	}
}

func (l *List) Remove2(n *ListNode) {
	node := l.Head
	if node == nil {
		return
	}

	if node == n {
		l.Head = node.Next
		return
	}

	for node != nil {
		if node.Next == n {
			node.Next = node.Next.Next
			// node = node.Next
			return
		}

		node = node.Next
	}
}
