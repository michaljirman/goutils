package main

import (
	"container/list"
	"fmt"
)

// Data structures - Linear
// A list is a sequence of elements. Each element can be connected to another with a link in a forward or backward direction.
func main(){
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element:= intList.Front(); element != nil; element=element.Next(){
		fmt.Println(element.Value.(int))
	}
}
