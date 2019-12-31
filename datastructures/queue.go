//+build disable

package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(n int) {
	q.items = append(q.items, n)
}

func (q *Queue) Dequeue() int {
	qLen := len(q.items)
	if qLen == 0 {
		return -1
	}
	item, items := q.items[0], q.items[1:]
	q.items = items
	return item
}

func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
