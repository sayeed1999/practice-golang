package main

import "fmt"

// Closure Example in Go: -
//
// This program demonstrates the concept of closures in Go.
// A closure is a function that captures the variables from its surrounding context.
// In this example, we have an outer function that defines a closure function.
// The closure function can access and modify the variables defined in the outer function,
// as well as variables defined in the global scope.
// The closure function is returned from the outer function and can be called multiple times,
// each time retaining the state of the variables it captured..

// Stack memory vs Heap memory: -
// Stack frame gets automatic cleanup, but heap needs Garbage Collector (GC) to cleanup..
// Stack memory is fixed, but heap memory is dynamic, can increase over necessity.
// That's why you see 'Stack Overflow' error!

const a = 10 // compile time creation

var p = 100 // runtime creation

// Go runs init() before any function is invoked!!
func init() {
	fmt.Println("=== learning closure ===")
}

// Note: -
// Once main func is finished executing, stack memory is totally cleaned up,
// but we have no control over heap memory.
// The GC might clean up immediately or much later after the main func is finished running!

func main() {
	showFn := outer()
	showFn() // calling the closure function
	// prints 210
	showFn() // calling the closure function again
	// prints 320
	showFn() // calling the closure function again
	// prints 430

	showFn2 := outer() // will have a fresh copy of the closure
	showFn2()          // prints 210 again
	showFn2()          // prints 320
	showFn2()          // prints 430
}

// Note: One

func outer() func() {
	money := 100 // Go runtime does ESCAPE analysis and moves the money variable to
	// 'heap memory' as this variable will be needed after the outer function returns.
	age := 30

	fmt.Println("Age:", age)

	// Note 1: When showFn is initialized, Go runtime knows money will be needed later,
	// so it does ESCAPE analysis & send it (money) to 'heap' memory.
	showFn := func() {
		money = money + a + p // using the variables from outer function & global scope
		fmt.Println(money)
	}

	return showFn // returning the closure function
}

// go run main.go => compile it => main (binary file) => ./main (run it automatically & clean up after execution)
// go build main.go =>compile it => main (binary file) => then you can run it with ./main manually
