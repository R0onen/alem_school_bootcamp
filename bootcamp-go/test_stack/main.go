package main

import (
	"bootcamp/stack"
	"fmt"
)

func main() {
	stack := stack.NewStack()
	stack.Push(10)
	stack.Push(20)
	fmt.Println(stack.Peek()) // 20
	stack.Pop()
	fmt.Println(stack.Peek()) // 10
}
