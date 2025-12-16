package main

import (
	"fmt"
	"time"
)

// go routines are function that leave the main thread and run in background then come back and join the main thread
// when main function is finished or are ready to return any value

// go routies do not stop the program flow and are non blocking ( async await promises )

func main() {

	go sayHello() // say hello takes one second to execute

	go printLetters()
	go printNumbers()

	// err := go doWork()
	var err error

	go func() {
		err = doWork() // running function in imediatly invoking anonymous func
	}()

	time.Sleep(2 * time.Second)

	if err != nil { // need to wait for err to return
		fmt.Println("error: ", err)
	}

}

func sayHello() {
	time.Sleep(time.Second * 1)
	fmt.Println("Hello from go routine")
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("Alphabet: ", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func printNumbers() {
	for i := range 5 {
		fmt.Println("Number:   ", i, time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(2 * time.Second)
	return fmt.Errorf("this is an error")
}
