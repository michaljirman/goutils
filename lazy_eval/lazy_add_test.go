package lazy_eval

import "testing"

var expensiveFcnCounter int

func expensiveFcn() int{
	expensiveFcnCounter++
	return 23
}

func TestLazyAdd(t *testing.T){
	n := Make(expensiveFcn)
	t.Logf("%d", n())
	t.Logf("%d", n()+42)
	if n()+42 != 65 {
		t.Fail()
	}
	if expensiveFcnCounter != 1 {
		t.Errorf("expensive fcn wasn't called only once, number of calls %d", expensiveFcnCounter)
	}
}
