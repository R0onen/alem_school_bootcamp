package list

func (l *List) PushBackLists(lists ...*List) {
	current := l.Head
	if l.Head == nil {
		return
	}

	for current.Next != nil {
		current = current.Next
	}

	for _, k := range lists {
		current.Next = k.Head
		for k.Head.Next != nil {
			k.Head = k.Head.Next
		}
		current = k.Head

	}
}
