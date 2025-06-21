package main

import "fmt"

func main() {

	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println(result)

	result = a - b
	fmt.Println(result)

	result = a * b
	fmt.Println(result)

	result = a / b
	fmt.Println(result)

	result = a % b
	fmt.Println(result)

	var number uint8 = 0
	number--
	fmt.Println(number)

	// overflow or underflow is dangerous in go because it wraps around instead of giving up an error
	// need to handle it and make functions like safe computation

}
