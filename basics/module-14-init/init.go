package main

import "fmt"

func init() {
	fmt.Println("Initializing the first init function")
}

func init() {
	fmt.Println("Initializing the second init function")
}

func main() {
	// init functions run sequentially
	// main funciton is executed after init functions

	// use case ->
	// seting up task like global variables or constants , read config files or env variables , open data base connections or schema migrations
	// best practices ->
	// avoid side effects , initialization order to be kept in mind  , documentation
	fmt.Println("Initializing the main function")
}
