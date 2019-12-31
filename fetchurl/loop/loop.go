package main

import "fmt"

func main() {
	data := make([]int, 4)

	loopData := func(handledData chan<- int) {
		defer close(handledData)
		for i := range data {
			handledData <- data[i]
		}
	}

	handledData := make(chan int)
	go loopData(handledData)
	for num := range handledData {
		fmt.Println(num)
	}

	//alternative
	//orDone := func(done, c <-chan interface{}) <-chan interface{} {
	//	valStream := make(chan interface{})
	//	go func() {
	//		defer close(valStream)
	//		for {
	//			select {
	//			case <-done:
	//				return
	//			case v, ok := <-c:
	//				if ok == false {
	//					return
	//				}
	//				select {
	//				case valStream <- v:
	//				case <-done:
	//				}
	//			}
	//		}
	//	}()
	//	return valStream
	//}
	//
	//done := make(chan interface{})
	////defer close(done)
	//handledData := make(chan interface{})
	//
	//data := make([]int, 4)
	//loopData := func(handledData chan<- interface{}) {
	//	defer close(handledData)
	//	for i := range data {
	//		handledData <- data[i]
	//	}
	//}
	//go loopData(handledData)
	//for num := range orDone(done, handledData) {
	//	fmt.Println(num)
	//}
	//close(done)

}
