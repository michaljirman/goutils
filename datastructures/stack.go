//+build disable

package main

import "fmt"

// Stack represents stack structure
type Stack struct {
	items []int
}

// Push adds item to the top of stack
func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

// Pop removes item from the top of stack
func (s *Stack) Pop() int {
	curLen := len(s.items)
	if curLen == 0 {
		return -1
	}
	item, items := s.items[curLen-1], s.items[:curLen-1]
	s.items = items
	return item
}

func main() {

	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack.Pop())

}
