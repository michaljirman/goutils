package function_scheduler_evaluator

import "testing"

//func TestFcnWithArguments(t *testing.T)  {
//	called := false
//	fnc1 := func(a... int) int {
//		called = true
//		return 0
//	}
//	le := LazyEvalution{}
//	le.FcnWithArguments(fnc1, 1, 2)
//	if called == true {
//		t.Fail()
//	}
//}
//
//func TestFcnWithArguments2(t *testing.T) {
//	fnc1 := func(a... int) int {
//		sum := 0
//		for _, n := range a {
//			sum += n
//		}
//		return sum
//	}
//	le := LazyEvalution{}
//	le.FcnWithArguments(fnc1, 2, 2)
//	le.FcnWithArguments(fnc1, 5, 7)
//	result := le.Evaluate()
//	if len(result) != 2 {
//		t.Errorf("unexpected result %+v", result)
//	}
//
//}


func TestFunctionEvaluator(t *testing.T) {
	fnc1 := func(a... int) int {
		sum := 0
		for _, n := range a {
			sum += n
		}
		return sum
	}
	fe := FunctionEvaluator{}
	fe.Add(fnc1, 2, 2)
	fe.Add(fnc1, 5, 7)
	result := fe.Evaluate()
	if len(result) != 2 {
		t.Errorf("unexpected result %+v", result)
	}
	if result[0] != 4 {
		t.Errorf("unexpected result %+v", result)
	}
	if result[1] != 12 {
		t.Errorf("unexpected result %+v", result)
	}
}