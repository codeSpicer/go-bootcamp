package main

import "fmt"

func main() {

	message := "hello world"

	for i, v := range message {
		fmt.Print(i, v, "     ") // prints unicodes
		fmt.Printf("Index: %v, Rune: %c\n", i, v)
	}

	fmt.Println("\n--- Working with Slices ---")

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", numbers)

	// When using range with a slice, 'v' is a copy of the element's value,
	// not a reference to the original element.
	// Therefore, modifying 'v' inside the loop will not affect the original slice.
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
		v = v * 2 // This modifies the copy 'v', not the element in 'numbers'
	}

	fmt.Println("Slice after ranging (values not modified by 'v = v * 2'):", numbers)

	// To modify the original slice elements using range, you need to
	// access them via their index:
	fmt.Println("\n--- Modifying Slice Elements ---")
	for i := range numbers {
		numbers[i] = numbers[i] * 2 // This modifies the original slice element
	}

	fmt.Println("Slice after modifying via index:", numbers)

}
