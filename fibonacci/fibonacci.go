package fibonacci

func fibonacci(n int, c chan int){
	x, y := 0, 1
	for i:=0; i<n; i++{
		c <-x
		x, y = y, x+y
	}
	close(c)
}

func Fibonacci(n int) []int{
	var fibSeq []int
	c:= make(chan int, n)
	go fibonacci(n, c)
	for i := range c {
		fibSeq = append(fibSeq, i)
	}
	return fibSeq
}

