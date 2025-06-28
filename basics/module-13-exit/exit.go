package main

import (
	"fmt"
	"os"
)

func main() {

	// os.exit is immediate exit without any cleanup , defer panic recover are ignored
	// use case is -> any critical error that cannot be fixed or recovered
	// zero is for sucess and non zero exit codes are for errors

	defer fmt.Println("defered statement")
	fmt.Println("starting the function")

	// exit with status code 1
	os.Exit(1)

}
