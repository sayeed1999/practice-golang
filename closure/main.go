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

const a = 10 // compile time creation

var p = 100 // runtime creation

// go run main.go =>

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

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age:", age)

	showFn := func() {
		money = money + a + p // using the variables from outer function & global scope
		fmt.Println(money)
	}

	return showFn // returning the closure function
}

// go run main.go => compile it => main (binary file) => ./main (run it automatically & clean up after execution)
// go build main.go =>compile it => main (binary file) => then you can run it with ./main manually
