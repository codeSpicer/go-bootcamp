package main

import (
	"fmt"
)

func main() {
	num := 42
	str := "GoLang"
	char := 'G'
	float := 3.14159
	boolean := true
	arr := []int{1, 2, 3}
	m := map[string]int{"one": 1, "two": 2}

	fmt.Println("Go Formatting Verbs Examples:")

	// %v - value in default format
	fmt.Printf("%%v (value): %v\n", arr)

	// %d - decimal integer
	fmt.Printf("%%d (decimal): %d\n", num)

	// %f - floating point
	fmt.Printf("%%f (float): %f\n", float)

	// %s - string
	fmt.Printf("%%s (string): %s\n", str)

	// %c - character
	fmt.Printf("%%c (character): %c\n", char)

	// %t - boolean
	fmt.Printf("%%t (boolean): %t\n", boolean)

	// %T - type of the value
	fmt.Printf("%%T (type): %T\n", arr)

	// %#v - Go-syntax representation of the value
	fmt.Printf("%%#v (Go-syntax): %#v\n", m)

	// %p - pointer address
	fmt.Printf("%%p (pointer): %p\n", &num)

	// %q - quoted string or character
	fmt.Printf("%%q (quoted string): %q\n", str)
	fmt.Printf("%%q (quoted char): %q\n", char)
}
