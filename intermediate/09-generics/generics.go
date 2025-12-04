package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {

	x, y := 1, 2
	x, y = swap(x, y)

	fmt.Println(x, y)

	x1, y1 := "akshat", "kumar"
	x1, y1 = swap(x1, y1)
	fmt.Println(x1, y1)
	
}
