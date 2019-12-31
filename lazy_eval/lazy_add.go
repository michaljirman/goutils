package lazy_eval

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func Make(f func() int) LazyInt{
	var v int
	var once sync.Once
	return func() int{
		once.Do(func() {
			fmt.Println("calling f()")
			v = f()
		})
		return v
	}
}
