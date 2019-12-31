package main

import (
	"fmt"
	"math/rand"
	"time"
)

//The quick sort algorithm falls under the divide and conquer class of algorithms,
// where we break (divide) a problem into smaller chunks that are much simpler to solve (conquer).
// In this case, an unsorted array is broken into sub-arrays that are partially sorted,
// until all elements in the list are in the right position,
// by which time our unsorted list will have become sorted.

//Although the worst case time complexity of QuickSort is O(n2)
// which is more than many other sorting algorithms like Merge Sort and Heap Sort,
// QuickSort is faster in practice, because its inner loop can be efficiently
// implemented on most architectures, and in most real-world data.
// QuickSort can be implemented in different ways by changing the choice of pivot,
// so that the worst case rarely occurs for a given type of data.
// However, merge sort is generally considered better when data is huge and stored in external storage.
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	fmt.Println(pivot, left, right)
	a[pivot], a[right] = a[right], a[pivot]
	fmt.Println(a)
	for i, _ := range a {
		fmt.Printf("%d < %d\n", a[i], a[right])
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
			fmt.Println(a)
		}
	}
	//http://www.golangprograms.com/data-structure-and-algorithms.html
	fmt.Println(a)
	a[left], a[right] = a[right], a[left]
	quicksort(a[:left])
	quicksort(a[left+1:])
	return a
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	//r := rand.Int()
	//fmt.Println(r)
	//fmt.Println(r % 20)


	//slice := generateSlice(20)
	slice := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

