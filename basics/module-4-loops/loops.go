package main

import (
	"fmt"
)

func main() {
	count := 6
	// Traditional for loop with counter
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	// iterate over a collection
	numbers := []int{1, 2, 3, 4, 5, 6}

	// Range loop over integer (0 to 9)
	for i := range 10 {
		fmt.Println(i)
	}

	// Range loop with index and value
	for index, value := range numbers {
		fmt.Printf("Index %d, value %d\n", index, value)
	}

	// For loop with continue and break statements
	for i := 0; i <= 6; i++ {
		if i%2 == 1 {
			continue
		} else {
			fmt.Printf("even number is: %v\n", i)
		}
		if i == 4 {
			break
		}
	}

	// While-style loop (condition-based)
	j := 0
	for j <= 5 {
		fmt.Println("iteration ", j, " ")
		j++
	}

}
