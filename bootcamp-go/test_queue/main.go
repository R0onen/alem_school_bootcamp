package main

import (
	"bootcamp/queue"
	"fmt"
)

func main() {
	queue := queue.NewQueue()
	queue.Push(10)
	queue.Push(20)
	fmt.Println(queue.Peek()) // 10
	fmt.Println(queue.Peek()) // 10
	queue.Pop()
	fmt.Println(queue.Peek()) // 20
}
