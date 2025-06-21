package main

import "fmt"

func main() {

	// Array declaration syntax: var arrayName [size] elementType
	var array [5]int
	fmt.Println("before operations", array)
	array[0] = 45 // Assign value to specific index
	fmt.Println("after change", array)

	// Array initialization with values
	fruits := [3]string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// Arrays are copied by value, not by reference
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[1] = 100 // This won't affect the original array
	fmt.Println("original ", originalArray)
	fmt.Println("copied ", copiedArray)

	// Iterating over array using range
	for i, v := range fruits {
		fmt.Printf("Index: %d , Value: %v\n", i, v)
	}

	// Dynamic arrays (slices) - more flexible than fixed arrays
	dynamicArray := []int{}
	for i := 0; i < 10; i++ {
		dynamicArray = append(dynamicArray, i) // Add elements dynamically
	}
	fmt.Println("dynamic array :", dynamicArray, " capacity of array ", cap(dynamicArray))

	// Multidimensional array (matrix)
	matrix := [4][3]int{}

	// Pointer to array
	var matrix2 *[4][3]int
	matrix2 = &matrix

	// Initialize matrix with values
	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			matrix[i][j] = i*j + 1
		}
	}
	fmt.Println(matrix)
	fmt.Println(matrix2)

	// Blank identifier (_) to ignore unwanted return values
	_, b := someFunction()
	_ = b // Ignore the second return value
}

// Function that returns multiple values
func someFunction() (int, int) {
	return 1, 2
}
