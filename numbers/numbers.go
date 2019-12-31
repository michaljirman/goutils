package numbers

func sum(numbers []int, c chan int){
	sum := 0
	for _, n:= range numbers{
		sum += n
	}
	c <- sum
}

func SumNumbers(numbers []int) int{
	c:= make(chan int)
	go sum(numbers[:len(numbers)/2], c)
	go sum(numbers[len(numbers)/2:], c)
	return <-c+<-c
}

