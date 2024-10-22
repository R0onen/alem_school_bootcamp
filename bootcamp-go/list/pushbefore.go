package list

type ListNode struct {
	Value interface{}
	Prev  *ListNode
	Next  *ListNode
}

type List struct {
	Head, Tail *ListNode
}

func (l *List) PushBefore(n *ListNode, v interface{}) {
	newNode := &ListNode{Prev: nil, Next: nil, Value: v}
	if l.Head == nil {
		l.Head = newNode
	}
	if n == l.Head {
		newNode.Next = n
		n.Prev = newNode
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != n {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
}
