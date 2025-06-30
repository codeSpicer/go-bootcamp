package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// A pointer is a variable that stores the memory address of another variable.
	// Pointer types are declared with an asterisk: var ptr *int

	// Initialization
	var intPtr *int
	number := 67
	intPtr = &number // intPtr now holds the address of number

	fmt.Printf("number is stored at memory address: %p, value at intPtr: %d\n", intPtr, *intPtr) // Dereferencing

	// Pointers with slices
	nums := []int{1, 2, 3, 4, 5}
	slicePtr := &nums
	fmt.Printf("slicePtr points to address: %p, value: %v\n", slicePtr, *slicePtr)

	// Nil pointer
	var nilPtr *int
	fmt.Printf("nilPtr: %v\n", nilPtr)
	if nilPtr == nil {
		fmt.Println("nilPtr is nil")
	}

	// Passing pointer to a function
	value := 3
	fmt.Println("Value before modifyValue:", value)
	modifyValue(&value)
	fmt.Println("Value after modifyValue:", value)

	// Pointer to struct example

	alice := Person{name: "Alice", age: 30}
	personPtr := &alice
	fmt.Printf("personPtr points to: %+v\n", *personPtr)
	updateAge(personPtr, 35)
	fmt.Printf("After updateAge, person: %+v\n", alice)

	// Note: There are also unsafe pointers, but they are rarely needed in typical Go code.
}

func modifyValue(ptr *int) {
	// Increment the value pointed to by ptr
	(*ptr)++
}

// updateAge updates the age of a Person via pointer
func updateAge(p *Person, newAge int) {
	p.age = newAge
}
