package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	// Using fmt.Scan (input separated by space, does not consume newline)
	fmt.Print("Enter your name and age (separated by space): ")
	fmt.Scan(&name, &age) // scan needs pointer of the variable to write input directly into the variable
	fmt.Println("Scan -> Name:", name, "Age:", age)

	// Using fmt.Scanln (input separated by space, stops at newline)
	var city string
	fmt.Print("Enter your city: ")
	fmt.Scanln(&city)
	fmt.Println("Scanln -> City:", city)

	// Error formatting example using fmt.Errorf
	// fmt.Errorf returns an error with a formatted message
	err := fmt.Errorf("user %s of age %d from %s encountered an error", name, age, city)
	fmt.Println("Formatted error:", err)

	// String formatting example using fmt.Sprintf
	// fmt.Sprintf returns a formatted string (does not print)
	info := fmt.Sprintf("User info: Name=%s, Age=%d, City=%s", name, age, city)
	fmt.Println("Sprintf ->", info)
}
