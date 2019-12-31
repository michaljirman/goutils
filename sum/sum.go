package sum

import (
	"fmt"
	"sync"
)

type SafeSum struct {
	wg *sync.WaitGroup
	mux sync.Mutex
	Sum int
}

func (s *SafeSum) Add(val int){
	s.mux.Lock()
	defer s.wg.Done()
	defer s.mux.Unlock()
	s.Sum += val
}

func ConcurrentSum(n []int) int{
	wg := &sync.WaitGroup{}
	wg.Add(len(n))
	s := &SafeSum{wg: wg}
	for _, n := range n {
		go s.Add(n)
	}
	wg.Wait()
	fmt.Println(s.Sum)
	return s.Sum
}