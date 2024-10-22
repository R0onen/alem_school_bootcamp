package main

import (
	"bootcamp/list"
	"fmt"
)

func main() {
	l := list.NewList()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.PushBack(40)

	l.RemoveIf(func(n *list.ListNode) bool {
		return n.Value.(int) > 20
	})

	l.ForEach(func(n *list.ListNode) {
		fmt.Println(n.Value)
	})
	// Output:
	// 10
	// 20
}
