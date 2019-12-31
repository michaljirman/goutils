package main

import (
	"container/heap"
	"fmt"
)

// Data structures - Linear
// A heap is data structure that is based on the heap property. The heap data structure is used in selection, graph,
// and k-way merge algorithms. Operations such as finding, merging, insertion, key changes, and deleting are performed on heaps.
// According to the heap order (maximum heap) property, the value stored at each node is greater than or equal to its children.

type IntegerHeap []int

// IntegerHeap method - gets the length of integerHeap
func (iHeap IntegerHeap) Len() int {
	return len(iHeap)
}
// IntegerHeap method - checks if element of i index is less than j index
func (iHeap IntegerHeap) Less(i, j int) bool {return iHeap[i] < iHeap[j]}

// IntegerHeap method - swaps the element of i to j index
func (iHeap IntegerHeap) Swap(i, j int) {iHeap[i], iHeap[j] = iHeap[j], iHeap[i]}

// IntegerHeap method - pushes the item
func (iHeap *IntegerHeap) Push(heapIntf interface{}){
	*iHeap = append(*iHeap, heapIntf.(int))
}

// IntegerHeap method - pops the item from the heap
func (iHeap *IntegerHeap) Pop() interface{}{
	var n int
	var x1 int
	var previous = *iHeap
	n = len(previous)
	x1 = previous[n-1]
	*iHeap = previous[0 : n-1]
	return x1
}

// main method
func main(){
	var iHeap = &IntegerHeap{2, 4, 1}

	heap.Init(iHeap)
	heap.Push(iHeap,3)
	fmt.Printf("minimum: %d\n", (*iHeap)[0])
	for iHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(iHeap))
	}
}