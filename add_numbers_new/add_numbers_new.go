package add_numbers_new

import (
	"fmt"
	"math/rand"
	"time"
)

var numbers [] int

func randInt() int{
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	time.Sleep()
	return r1.Intn(100)
}

func add(c chan int){
	n:= randInt()
	//fmt.Println(n)
	c <- n
}

func AddNumbers(){
	total := 200
	c := make(chan int, total)
	for i:=0; i<total; i++{
		go add(c)
	}

	for i:=0; i<total; i++{
		numbers = append(numbers, <-c)
	}
	fmt.Println(numbers)
}