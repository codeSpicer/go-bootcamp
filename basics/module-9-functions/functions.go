// Package main provides examples of various function types and their usage in Go.
package main

import (
	"fmt"
	"reflect"
)

// function is the entry point of the program.
func FunctionsWithSingleReturn() {

	// Function declaration syntax:
	// func funcName(parameters list) returnType {
	// 	// code block
	// 	// return value
	// }

	// Public functions in Go start with a capital letter (e.g., fmt.Println).
	// Private functions start with a lowercase letter.

	// Calling a simple function and printing its result.
	sum := add(4, 5)
	fmt.Println("Sum of 4 and 5:", sum)

	// Anonymous function: a function without a name, stored in a variable and then called.
	greet := func() {
		fmt.Println("Hello from an anonymous function!")
	}
	greet()

	// Functions as First-Class Citizens:
	// Functions can be assigned to variables, passed as arguments to other functions, and returned from functions.
	operation := add // 'add' function assigned to a variable 'operation'.

	// Printing the type of the 'operation' variable, which is a function type.
	fmt.Println("Type of 'operation' variable:", reflect.TypeOf(operation))

	// Calling the function through the variable.
	fmt.Println("Result of 'operation(3, 7)':", operation(3, 7))

	// Passing a function as an argument to another function.
	fmt.Println("Result of 'funcCaller(operation, 20, 30)':", funcCaller(operation, 20, 30))

	// Function returning another function (closure).
	// 'multiplier' returns a new anonymous function that multiplies its input by 'factor'.
	multiplyByFour := multiplier(4)
	fmt.Println("Result of 'multiplyByFour(10)':", multiplyByFour(10))
}

// add takes two integers and returns their sum.
// Modifications to the parameters within the function do not affect the original arguments.
func add(a, b int) int {
	return a + b
}

// funcCaller takes a function (x) that accepts two integers and returns an integer,
// along with two integer arguments (a, b).
// It then calls the passed function 'x' with 'a' and 'b' and returns its result.
func funcCaller(x func(int, int) int, a, b int) int {
	return x(a, b) // Calls the passed function and returns its output.
}

// multiplier is a higher-order function that returns a new function (a closure).
// The returned function takes an integer 'x' and multiplies it by the 'factor'
// provided when 'multiplier' was called.
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
