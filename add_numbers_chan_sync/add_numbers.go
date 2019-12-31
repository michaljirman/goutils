package add_numbers_chan_sync

import "fmt"

var numbers []int

func AddNumbers(){
	ch := make(chan int, 10)
	go func(ch chan int) {
			for i:=0; i<10; i++{
				ch <- i
			}
			close(ch)
	}(ch)
	for n := range ch{
		numbers = append(numbers, n)
	}
	fmt.Printf("%+v", numbers)
}