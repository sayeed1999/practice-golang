package main

import (
	"fmt"
	"time"
)

// The root cause: The undiscovered panic in any goroutine terminates the entire program immediately,
// and deferred functions only run in the goroutine that is unwinding, not that is the main goroutine.

func main() {
	fmt.Println("main goroutine: started main() normally")
	defer fmt.Println("main goroutine: ending main() normally") // never printed!

	go func() {
		defer fmt.Println("separate goroutine: ending main() normally") // printed!
		panic("separate goroutine: something went wrong")
	}()

	time.Sleep(5 * time.Second)
}
