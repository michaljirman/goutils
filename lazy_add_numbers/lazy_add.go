package lazy_add_numbers

import "sync"

type LazyAddFunction func(a, b int) int

func Make(f LazyAddFunction) LazyAddFunction{
	var v int
	var once sync.Once
	return func(a, b int) int {
		once.Do(func() {
			v = f(a, b)
		})
		return v
	}
}
