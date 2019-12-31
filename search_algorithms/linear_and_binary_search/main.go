package main

import "fmt"

//This technique pass over the list of elements, by using the index to move from the beginning of the list to the end.
// Each element is examined and if it does not match the search item, the next item is examined.
// By hopping from one item to its next, the list is passed over sequentially.
func linearSearch(dataList []int, key int) bool{
	for _, item := range dataList{
		if item == key {
			return true
		}
	}
	return false
}

// A binary search is a search strategy used to find elements within a list by consistently reducing the amount of data
// to be searched and thereby increasing the rate at which the search term is found. To use a binary search algorithm,
// the list to be operated on must have already been sorted.
func binarySearch(needle int, haystack[]int) bool {
	low:= 0
	high:= len(haystack)-1
	for low<=high{
		fmt.Println(low, high)
		median := (low + high)/2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median -1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	return true
}

func main()  {
	items:= []int{95,78,46,58,45,86,99,251,320}
	fmt.Println(linearSearch(items, 58))

	items = []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}
