// We want to build golang library package which will allow to pass functions and arguments to it.
// It will not evaluate the function right away but once developer decides.
package function_scheduler_evaluator

type fcnType func(a... int) int

type FunctionEvaluate struct {
	fcn fcnType
	args []int
}

type FunctionEvaluator struct {
	fcnEvaluates []FunctionEvaluate
}

func (e *FunctionEvaluator) Add(fcn fcnType, n...int){
	e.fcnEvaluates = append(e.fcnEvaluates, FunctionEvaluate{fcn: fcn, args:n})
}

func (e *FunctionEvaluator) Evaluate()[]int{
	var res []int
	for _, fe := range e.fcnEvaluates{
		res = append(res, fe.fcn(fe.args...))
	}
	return res
}

//type LazyEvalution struct {
//	fcn []add
//	args [][]int
//}
//
//func (e *LazyEvalution)FcnWithArguments(fcn add, n ...int){
//	e.args = append(e.args, n)
//	e.fcn = append(e.fcn, fcn)
//}
//
//func (e *LazyEvalution)Evaluate() []int{
//	var result []int
//	for i, f := range e.fcn {
//		result = append(result, f(e.args[i]...))
//	}
//	return result
//}