package main

import (
	"fmt"

	"github.com/bradford-hamilton/range-func-playground/stack"
)

// GOEXPERIMENT=rangefunc
func main() {
	stack1 := stack.New(0, 1, 2, 3, 4, 5, 6, 7)
	stack1.Push(8)
	stack1.Push(9)
	stack1.Push(10)

	for v := range stack1.TopDown() {
		fmt.Printf("%+v ", v)
	}

	fmt.Println()

	// --------------------------------------------------- //

	stack2 := stack.New(0, 1, 2, 3, 4, 5, 6, 7)
	stack2.Push(8)
	stack2.Push(9)
	stack2.Push(10)

	for v := range stack2.BottomUp() {
		fmt.Printf("%+v ", v)
	}

	fmt.Println()

	item1, ok1 := stack2.Pop()
	item2, ok2 := stack2.Pop()
	item3, ok3 := stack2.Pop()

	fmt.Println(item1, ok1)
	fmt.Println(item2, ok2)
	fmt.Println(item3, ok3)

	for v := range stack2.BottomUp() {
		fmt.Printf("%+v ", v)
	}

	fmt.Println()
}
