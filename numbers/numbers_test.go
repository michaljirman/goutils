package numbers

import "testing"

func TestSumNumbers(t *testing.T){
	numbers:= []int{7, 2, 8, -9, 4, 0}
	sum:= SumNumbers(numbers)
	if sum != 12 {
		t.Log("invalid sum")
		t.Fail()
	}
}