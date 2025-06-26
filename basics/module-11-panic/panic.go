// A panic is a built-in function that stops the ordinary flow of control and begins panicking.
// When the function F calls panic, the execution of F stops, any deferred functions in F are executed normally,
// and then F returns to its caller. To the caller, F then behaves like a call to panic.
// The process continues up the stack until all functions in the current goroutine have returned,
// at which point the program crashes. Panics can be recovered by a call to recover.
package main

import "fmt"

func main() {
	fmt.Println("Calling process with a valid number.")
	process(4)
	fmt.Println("---")

	fmt.Println("Calling process with a number that will cause a panic.")
	// The program will exit after this function call panics.
	process(-1)

	// This code is unreachable because the program terminates during the panic.
	fmt.Println("Program finished successfully.")
}

// process function checks if the input number is positive.
// It panics if the number is negative.
func process(a int) {
	// This deferred call will execute regardless of whether the function panics or returns normally.
	// This is useful for cleanup activities.
	defer fmt.Println("Cleanup: processed number", a)
	if a < 0 {
		// panic is used here to signal an unexpected and unrecoverable error condition.
		// This might indicate a bug in the code that calls this function.
		panic("input number must be positive")
	}
	fmt.Println("Processing valid input:", a)
}
