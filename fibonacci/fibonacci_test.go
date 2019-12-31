package fibonacci

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFibonacci(t *testing.T){
	fibSeq := Fibonacci(10)
	expectedFibSeq:= []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	assert.Equal(t, fibSeq, expectedFibSeq, "unexpected result for fib sequence")
}
