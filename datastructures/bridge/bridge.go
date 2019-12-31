package main

import "fmt"

// Bridge - structural design patterns
// Bridge decouples the implementation from the abstraction. The abstract base class can be subclassed to provide
// different implementations and allow implementation details to be modified easily.
// The interface, which is a bridge, helps in making the functionality of concrete classes independent from the interface
// implementer classes. The bridge patterns allow the implementation details to change at runtime.

type IDrawShape interface {
	drawShape(x[5] float32, y[5] float32)
}
type DrawShape struct {}

func (drawShape DrawShape) drawShape(x[5] float32, y[5] float32){
	fmt.Println("Drawing shape")
}

type IContour interface {
	drawContour(x[5] float32, y[5] float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x[5] float32
	y[5] float32
	shape DrawShape
	factor int
}

func (contour DrawContour) drawContour(x[5] float32, y[5] float32){
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}
func (contour DrawContour) resizeByFactor(factor int){
	contour.factor = factor
}

func main(){
	var x = [5]float32{1,2,3,4,5}
	var y = [5]float32{1,2,3,4,5}
	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}

