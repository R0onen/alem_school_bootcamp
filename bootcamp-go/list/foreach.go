package list

func (l *List) ForEach(fn func(n *ListNode)) {
	for current := l.Head; current != nil; current = current.Next {
		fn(current)
	}
}
