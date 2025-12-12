package main

import (
	"errors"
	"fmt"
	// "math"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// Return 0 and a non-nil error value
		return 0, errors.New("cannot divide by zero")
	}
	// Return the result and a nil error
	return a / b, nil
}

func basicErrors() {
	result, err := Divide(10, 0)

	if err != nil { // Check if an error occurred
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result) // Output: Result: 5

}
