package list

func (l *List) Sort(fn func(a *ListNode, b *ListNode) int) {
	if l.Head == nil || l.Head == l.Tail {
		return
	}

	l.Head = mergeSort(l.Head, fn)
}

func mergeSort(head *ListNode, fn func(a *ListNode, b *ListNode) int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := findMiddle(head)

	nextToMiddle := middle.Next
	middle.Next = nil

	left := mergeSort(head, fn)
	right := mergeSort(nextToMiddle, fn)

	return merge(left, right, fn)
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(left, right *ListNode, fn func(a *ListNode, b *ListNode) int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for left != nil && right != nil {
		if fn(left, right) <= 0 {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}

	return dummy.Next
}
