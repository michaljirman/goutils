#### Q. Mention some of the key features of Golang.
Ans: Fast(A performant language)
Good support for packages(I come from C++ background)
Goroutines instead of threads(with channels need not worry much about mutex, etc)
Very different way of implementing interfaces.
Inheritance is not explicit but implicit
Prefers composition over inheritance
Compiles fast
Easy to learn

####Q. Difference between parallelism, concurrency and multi threading?
Ans: Parallelism is executing multiple parts of a program parallely maybe on a multi core CPU.
E.g. cross product of a vector
Whereas concurrency is the way a program is structured. The program is composed of multiple independent parts which may or may not execute in parallel(e.g. on single core CPU). But can be made to work parallely if needed. It is composing a program of multiple independent components.
E.g. How mouse, keyboard, audio drivers work.

####Q. How is a thread different from go routine?
Ans: Goroutines are light weight.
OS threads mostly take around 2MB. So 1024 threads take 2GB RAM.
A Goroutine starts life with a small stack, typically 2 KB and stack of Goroutine is growable maybe even up to 1GB. This lets you create millions of go routines which will be difficult using threads.
In case of threads if it grows beyond 2MB might lead to stack overflow.
Another good thing about go routines is it
Goroutine lets you avoid locking hell.
Threads increase complexity and brittleness of the codebases. There can be latent deadlocks and race conditions, and it can become near impossible to reason about the code. In Go, if two goroutines need to share data, they can do so safely over a channel. Go handles all of the synchronization for you, and it’s much harder to run into things like deadlocks. The mantra here is don’t communicate by sharing memory, share memory by communicating.

####Q. How to create a non blocking channel?
Ans: Using default case in select case.
Here is an example.
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```


####Q. Can a sender close a channel?
Ans: Yes, sender can close the channel. Never the receiver should close a channel. Sending on a closed channel will cause panic.
v, ok := <-ch
Receiver can check whether the channel is closed using the second value. 'ok' being False indicates channel is closed.
Receiving from closed channel will return zero value(default value) immediately.

####Q. What is a nil channel?
Ans: Following is an example of nil channel:
var c1 chan int
//c1 = make(chan int)//This is commented
In this case buffer size of channel is 0. Reading/writing from/to such a channel will block forever.

####Q. How do we ensure that all goroutines have ended?
Ans: sync.WaitGroup() offers convenient way of achieving this.
var wg sync.WaitGroup
Before start of every goRoutine we need to do wg.Add(1) or do wg.Add(n) upfront if number of goroutines are known.
Do defer wg.Done() within each goroutine func. And let main wg.Wait().
If wg.Done() is called when count is 0, it leads to panic. E.g. if we forget to do wg.Add(1) somewhere.

####Q. Difference between function and method in golang

####Q. What all can act/or cannot act as receivers of methods?
Ans: Structs, pointers to structs, aliases of even built-in types such as 'type myIntType int'. Now myIntType can act as receiver.
Even functions can be receivers.
We can’t use the following things as receiver types:
methods- if we define method on object type it can’t be used as receiver type just like a normal function.
interfaces- In Go, interfaces are defining set of actions possible for the type. They are not defining actual implementation. Hence they can’t be used as receivers for methods, because methods are about implementation.

####Q. Difference between arrays and slices in Golang?
Ans: Slices are a wrapper around arrays in Golang.
Always prefer slices over arrays.
There are very few cases where an array will be beneficial.
One case is when the size of array is fixed (say to store IPv4 address).
Or say we need a slice inside a struct. In this case, we need to do memory allocation twice i.e. once for struct and then for slice. This will slow down the performance which is why we may use an array here. Maybe allocate it some additional memory. It will be a bit memory inefficient but performance will improve.
For passing around functions a slice will be better. Passing arrays over functions is very expensive since it does pass by value.
When we do not want functions to modify original copy we might use arrays. Even in that case there should be a way to pass slices.

####Q. What are runtime variables in Golang?
Ans: Here is the answer.

####Q. Difference between select and switch case?

####Q. How can race conditions be avoided?
Ans: Here is a good read.

####Q. Difference between interfaces of golang vs other languages

####Q. How is the compiler of Golang different than other language compilers

####Q. How do you end a goroutine

####Q. How does parent goroutine know if a goroutine is ended

####Q. Implement ping pong using channel.


