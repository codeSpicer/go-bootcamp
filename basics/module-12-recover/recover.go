package main

import "fmt"

func main() {
	// recover is used inside defer when panic happens
	// use case ->
	// recovery , cleanup , logging and reporting to fix production issue without stoping the system
	// always use recover with defer, better to log and report instead of silently recovering , should only be used for non recoverable errors

	process()
	fmt.Println("Returned from process")

	fmt.Println("\n--- Recovering from a function that returns a value ---")
	val := recoverAndReturnValue()
	fmt.Printf("Returned from recoverAndReturnValue: %q\n", val)
}

func process() {
	defer func() {
		// r := recover()
		// if r != nil {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()
	fmt.Println("start process")
	panic("something went wrong")
	fmt.Println("end process")
}

// recoverAndReturnValue demonstrates recovering from a panic in a function
// that is expected to return a value.
func recoverAndReturnValue() (s string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in recoverAndReturnValue:", r)
			// we can modify the named return value
			s = "this value was recovered"
		}
	}()

	fmt.Println("Inside recoverAndReturnValue, about to panic")
	panic("a panic happened")

	// This part of the code is never reached
	return "this is the normal return"
}
