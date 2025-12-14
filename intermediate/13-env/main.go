package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println(user, home)

	err := os.Setenv("FRUIT", "APPLE")

	if err != nil {
		fmt.Println("error setting env var", err)
	}

	for _, e := range os.Environ() {
		kvPair := strings.SplitN(e, "=", 2)
		fmt.Println(kvPair[0])
	}

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("error unsetting env var", err)
	}

}
