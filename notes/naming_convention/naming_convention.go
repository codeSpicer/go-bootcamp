package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// PascalCase
	// Structs , Interfaces , Enums
	// CalculateArea , UserInfo , NewHttpRequest

	// snake_case
	// variables , file_names
	// user_id , first_name , http_request

	// UPPERCASE
	// constants ( stand out and immutablitly is specified )
	// MATH_PI

	// camelCase
	// variables , identifiers

	// either use camel case or snake case across the application

	const MAXRETRIES = 5

	var employeeId = 1001

	fmt.Println("Employee Id is ", employeeId)
}
