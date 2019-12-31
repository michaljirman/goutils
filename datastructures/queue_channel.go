//+build disable

package main

import "fmt"

type Queue struct {
	items chan int
}

func (q *Queue) Enqueue(n int) {
	q.items <- n
}

func (q *Queue) Dequeue() int {
	return <-q.items
}

func main() {
	q := Queue{
		items: make(chan int, 16),
	}
	q.Enqueue(1)
	q.Enqueue(2)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
}
