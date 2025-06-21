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

	// SCOPE
	fmt.Println(globalName)

}

func printName() {
	firstName := "akshat"
	fmt.Println(firstName)
}
