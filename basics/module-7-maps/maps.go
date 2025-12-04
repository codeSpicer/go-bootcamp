package main

import (
	"fmt"
)

// This program demonstrates the usage of maps in Go.
// Maps are unordered collections of key-value pairs, similar to dictionaries in other languages.

func main() {

	// Declaration of a map:
	// var mapVariable map[keyType]valueType
	// To initialize a map, use the make function:
	// m = make(map[keyType]valueType)

	// You can also declare and initialize a map with values:
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	// Example: Creating a map with int keys and string values
	mp := make(map[int]string)

	// Adding key-value pairs to the map
	mp[1] = "one"
	mp[2] = "two"

	// Printing the entire map
	fmt.Println(mp) // Output: map[1:one 2:two] (order is not guaranteed)

	// Accessing a value by key
	fmt.Println(mp[1]) // Output: one

	// Iterating over a map using range
	for k, v := range mp {
		fmt.Println("key", k, " value", v)
	}

	// Accessing a non-existent key returns the zero value for the value type
	fmt.Println("value of mp[3] is :", mp[3]) // Output: "" (empty string, since value type is string)

	// Deleting a key-value pair from the map
	// delete(mp, 2)
	// fmt.Println(mp)

	// Clearing all entries from the map (Go 1.21+)
	// clear(mp)
	// fmt.Println(mp)

	// Checking if a key exists in the map
	value, ok := mp[2]     // ok is true if key exists, false otherwise
	fmt.Println(value, ok) // Output: two true

	_, exists := mp[3]  // Check for a key that doesn't exist
	fmt.Println(exists) // Output: false

	// Iterating over values only
	for _, v := range mp {
		fmt.Println("value in map", v)
	}

	// Checking if a map is nil (uninitialized)
	// Note: A map created with make is not nil, but an uninitialized map is nil
	// clear(mp)
	if mp == nil {
		fmt.Println("map is empty")
	}
	// You cannot add values to a nil map; use make to initialize it

	// Getting the number of key-value pairs in the map
	fmt.Println(len(mp)) // Output: 2

	// Multi-dimensional maps: maps whose values are also maps
	twoDmap := make(map[string]map[int]string)
	twoDmap["map1"] = mp // Assigning an existing map as a value

	fmt.Println(twoDmap) // Output: map[map1:map[1:one 2:two]]

	// maps in golang are used as sets
	numSet := map[int]struct{}{} // struct{}{} takes zero byte

	numSet[3] = struct{}{} // adding element

	_, exists = numSet[3]
	fmt.Println(exists)
	
	delete(numSet, 1)

	fmt.Println(len(numSet))

}
