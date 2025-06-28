package main

import "fmt"

func main() {
	// ========================================
	// CLOSURES IN GO
	// ========================================
	// A closure is a function that captures and retains references to variables
	// from its outer (enclosing) scope, even after the outer function has finished executing.
	//
	// Key characteristics:
	// - Functions that can access external variables
	// - Functions that can return other functions
	// - Maintain state between function calls
	// - Variables are captured by reference, not by value

	// ========================================
	// BASIC CLOSURE EXAMPLE
	// ========================================
	sequence := adder() // adder() returns a closure function

	// Each call to sequence() increments the captured variable 'i'
	fmt.Println(sequence()) // Output: 1
	fmt.Println(sequence()) // Output: 2
	fmt.Println(sequence()) // Output: 3

	// ========================================
	// PRACTICAL USE CASES
	// ========================================

	// 1. Stateful Functions (like our adder example)
	fmt.Println("\n--- Stateful Functions ---")
	counter := createCounter(10) // Start from 10
	fmt.Println(counter())       // 11
	fmt.Println(counter())       // 12

	// 2. Function Factories
	fmt.Println("\n--- Function Factories ---")
	multiplier := createMultiplier(5)
	fmt.Println("5 * 3 =", multiplier(3)) // 15
	fmt.Println("5 * 7 =", multiplier(7)) // 35

	// 3. Callback Functions
	fmt.Println("\n--- Callback Functions ---")
	numbers := []int{1, 2, 3, 4, 5}
	filtered := filter(numbers, func(n int) bool {
		return n%2 == 0 // Filter even numbers
	})
	fmt.Println("Even numbers:", filtered) // [2 4]

	// 4. Middleware Pattern
	fmt.Println("\n--- Middleware Pattern ---")
	logger := createLogger("API")
	logger("User login") // Output: [API] User login
	logger("Data fetch") // Output: [API] Data fetch
}

// ========================================
// CLOSURE IMPLEMENTATIONS
// ========================================

// adder() demonstrates a basic closure that maintains state
func adder() func() int {
	i := 0 // This variable is captured by the returned function
	fmt.Println("Initializing adder with i =", i)

	// Returns an anonymous function (closure) that captures variable 'i'
	return func() int {
		i++ // Modifies the captured variable
		fmt.Println("Incrementing i to:", i)
		return i
	}
}

// createCounter demonstrates parameterized closures
func createCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

// createMultiplier demonstrates function factories
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

// filter demonstrates closures as callbacks
func filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

// createLogger demonstrates closures for configuration
func createLogger(prefix string) func(string) {
	return func(message string) {
		fmt.Printf("[%s] %s\n", prefix, message)
	}
}

// ========================================
// BEST PRACTICES & CONSIDERATIONS
// ========================================

/*
BEST PRACTICES:
1. Keep closures simple and focused on a single responsibility
2. Use descriptive names for closure variables
3. Consider the lifetime of captured variables
4. Be mindful of memory usage with large captured data
5. Use closures for stateful operations, callbacks, and function factories

CONSIDERATIONS:
1. Memory Management: Captured variables stay in memory as long as the closure exists
2. Variable Capture: Variables are captured by reference, not by value
3. Goroutine Safety: Be careful when using closures with goroutines
4. Performance: Closures have minimal overhead but can accumulate memory
5. Testing: Closures can make unit testing more complex

COMMON USE CASES:
1. Iterators and generators
2. Middleware functions
3. Event handlers and callbacks
4. Configuration functions
5. State management in concurrent programs
6. Function currying and partial application
*/
