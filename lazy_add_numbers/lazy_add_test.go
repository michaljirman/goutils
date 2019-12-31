package lazy_add_numbers

import (
	"sync"
	"testing"
)

var addNumbersCounter int
func addNumbers(a, b int) int{
	addNumbersCounter++
	return a+b
}

func TestLazyAdd(t *testing.T){
	lazyAddNumbers := Make(addNumbers)
	var resultOne, resultTwo int
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		resultOne = lazyAddNumbers(10, 20)
	}()
	go func() {
		defer wg.Done()
		resultTwo = lazyAddNumbers(10, 20)
	}()
	wg.Wait()

	if addNumbersCounter != 1 {
		t.Errorf("expensive fcn wasn't called only once, number of calls %d", addNumbersCounter)
	}

	if resultOne != 30 {
		t.Log(resultOne)
		t.Fail()
	}

	if resultTwo != 30 {
		t.Log(resultTwo)
		t.Fail()
	}
}
