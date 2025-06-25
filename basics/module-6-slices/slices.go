package main

import (
	"fmt"
	"slices"
)

func sliceDeclarations() {

	// Slices in Go: dynamic, flexible views into arrays.
	// Declaration: var sliceName []ElementType

	var numbers []int             // Declares an empty slice of ints (nil slice)
	var numbers1 = []int{1, 2, 3} // Declares and initializes a slice with values

	numbers2 := []int{4, 5, 6, 7, 8, 9} // Shorthand declaration and initialization

	slice := make([]int, 5) // Creates a slice of ints with length 5 (all elements zero)

	slice2 := numbers2[1:5] // Creates a new slice from numbers2, elements at index 1 to 4 (slicing is half-open: [start:end))

	slice = append(slice, 10, 11, 12) // Appends elements to slice; returns a new slice (original may not change if capacity exceeded)

	sliceCopy := make([]int, len(slice)) // Creates a new slice with the same length as 'slice'
	copy(sliceCopy, slice)               // Copies elements from 'slice' into 'sliceCopy'

	for i, v := range slice2 {
		fmt.Println(i, v)
	}

	fmt.Println(slices.Equal(slice, slice2))

	// Note: var nilSlice []int // Declares a nil slice (no underlying array, length and capacity = 0)

	fmt.Println(numbers, numbers1, numbers2, slice, slice2, sliceCopy)

}

func main() {

	// sliceDeclarations()

	// 2D Slices (Matrix): Slices can contain other slices, creating multi-dimensional data structures.
	matrix := make([][]int, 5) // Creates a slice of 5 slices of ints

	// Populate the matrix with a triangular pattern
	for i := 0; i < 5; i++ {
		innerSize := i + 1
		matrix[i] = make([]int, innerSize) // Each inner slice has a different length
		for j := 0; j < innerSize; j++ {
			matrix[i][j] = i + j // Assign values based on indices
		}
	}

	fmt.Println(matrix) // Print the resulting 2D slice

	// Slice an existing slice
	slice := []int{4, 5, 6, 7, 8, 9}
	slice2 := slice[1:4] // slice2 gets elements from index 1 up to (but not including) 4 -> {5, 6, 7}

	// Length vs. Capacity
	// Length is the number of elements in the slice.
	// Capacity is the number of elements in the underlying array, from the start of the slice.
	fmt.Println(len(slice2), cap(slice2), slice2) // len: 3, cap: 5 (from index 1 of 'slice' to the end), slice2: [5 6 7]

}
