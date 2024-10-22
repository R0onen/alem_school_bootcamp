package list

func (l *List) PushAfter(n *ListNode, v interface{}) {
	newNode := &ListNode{Value: v}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		if current == n {
			break
		}
		current = current.Next
	}
	newNode.Prev = current
	newNode.Next = current.Next

	current.Next = newNode

	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
}
