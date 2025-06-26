package main

import "fmt"

func VariadictFunctions() {

	// func functionName(param1 type1 , param2 type2 , param3 ...type3) returnType {
	// code block
	// }

	fmt.Println(sum(2, 4, 3, 5, 62, 24))

	slice := []int{123, 1432, 245, 234, 25, 56, 6}

	fmt.Println(sum(slice...))

}

func sum(nums ...int) int { // ellipsis can only be used for the last param
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}
