package list

func (l *List) Find(fn func(v interface{}) bool) *ListNode {
	current := l.Head
	for current != nil {
		if fn(current.Value) {
			return current
		}
		current = current.Next
	}
	return nil
}
