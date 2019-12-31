package main

import "fmt"

// Adapter - structural design patterns
// Adapter provides a wrapper with an interface required by the API client
// to link incompatible types and act as a translator between the two types.
// The adapter uses the interface of a class to be a class with another compatible interface.
//
// The adapter pattern comprises the target, adaptee, adapter, and client:
// - target is the interface that the client calls and invokes methods on the adapter and adaptee
// - the client wats the incompatible interface implemented by the adapter
// - the adapter translates the incompatible interface of the adaptee into an interface the client wants
type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}
func (adapter Adapter) process(){
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

type Adaptee struct {
	adapterType int
}
func (adaptee Adaptee) convert(){
	fmt.Println("Adaptee convert method")
}

func main(){
	var processor IProcess = Adapter{}
	processor.process()
}

