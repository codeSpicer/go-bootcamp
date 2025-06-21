package main

import "fmt"

// need to use var globally
var globalName = "global"
var name = "globalname"

func main() {

	var age int
	var name string = "localname"

	var name1 = "jane"

	// type is infered when we use the := symbol ( can only be used inside functions / local variable init)
	count := 10
	lastName := "akshat"

	fmt.Println(age, name1, name, count, lastName)

	// Default value:
	// numeric types ->  0
	// boolean types -> false
	// string types -> ""
	// pointer , structs , slice , map  , functions -> nil

	// SCOPE is blocked scope in go
	fmt.Println(globalName)

	const pi = 3.14
	const gravity float32 = 9.8

	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
	)
}

func printName() {
	firstName := "akshat"
	fmt.Println(firstName)
}
