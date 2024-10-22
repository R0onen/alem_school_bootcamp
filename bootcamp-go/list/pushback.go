package list

func (l *List) PushBack(v interface{}) {
	newNode := &ListNode{Value: v}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
