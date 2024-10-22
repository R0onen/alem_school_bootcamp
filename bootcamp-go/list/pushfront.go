package list

func (l *List) PushFront(v interface{}) {
	newNode := &ListNode{Value: v}
	newNode.Next = l.Head
	l.Head = newNode
}
