package main

import "fmt"

// Data structures - Linear
// A tuple is a finite sorted list of elements. It is a data structure that groups data.
// Tuples are typically immutable sequential collections.


//gets the power series of integer `a` and returns tuple of square of `a` and cube of `a`
func powerSeries(a int)(int, int){
	return a*a, a*a*a
}

func powerSeriesWithNamedReturnParams(a int)(square int, cube int){
	square = a*a
	cube = a*a*a
	return
}

func main(){
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square ", square, "Cube ", cube)
}
