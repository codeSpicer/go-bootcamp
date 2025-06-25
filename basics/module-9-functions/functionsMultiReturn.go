package main

import (
	"fmt"
)

// FunctionsWithMultiReturn demonstrates various ways to use functions that return multiple values.
func FunctionsWithMultiReturn() {

	// Function returning two values (sum and difference)
	a, b := 10, 3
	sum, diff := sumAndDiff(a, b)
	fmt.Printf("sumAndDiff(%d, %d): sum = %d, diff = %d\n", a, b, sum, diff)

	// Function returning quotient and remainder
	quot, rem := divideAndRemainder(17, 5)
	fmt.Printf("divideAndRemainder(17, 5): quotient = %d, remainder = %d\n", quot, rem)

	// Function returning named values
	min, max := minMax(8, 15)
	fmt.Printf("minMax(8, 15): min = %d, max = %d\n", min, max)

	// Ignoring one of the returned values using blank identifier
	_, onlyMax := minMax(4, 9)
	fmt.Printf("minMax(4, 9): only max = %d (min ignored)\n", onlyMax)
}

// sumAndDiff takes two integers and returns their sum and difference.
func sumAndDiff(x, y int) (int, int) {
	return x + y, x - y
}

// divideAndRemainder returns the quotient and remainder of integer division.
func divideAndRemainder(a, b int) (int, int) {
	return a / b, a % b
}

// minMax returns the minimum and maximum of two integers using named return values.
func minMax(a, b int) (min int, max int) {
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return // returns named values
}
