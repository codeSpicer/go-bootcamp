package main

import "fmt"

func main() {

	//A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	// use cases ->
	// used to ensure connections to files and databases are closed after use
	// used to unlock mutexes even if a function panics
	// logging and tracing

	// avoid complex logic

	// defer is something like finally block for a function, it can do cleanup actions

	process()

	fmt.Println(eval(3))

	defer fmt.Println("world")
	fmt.Println("hello")
}

func process() {
	defer fmt.Println("first out")
	defer fmt.Println("last in")
}

func eval(a int) int {
	defer fmt.Println("defered:", a)
	a++
	fmt.Println("normal:", a)
	return a
}
