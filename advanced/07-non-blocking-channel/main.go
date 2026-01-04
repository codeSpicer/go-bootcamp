package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	// non blocking receive
	select {
	case msg := <-ch:
		fmt.Println("Received: ", msg)
	default:
		fmt.Println("No element present")

	}

	// non blocking send
	select {
	case ch <- 1:
		fmt.Println("sent data to channel")
	default:
		fmt.Println("Channel not ready to receive")
	}

	// non blocking operation in realtime system

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received, ", d)
			case <-quit:
				fmt.Println("Stopping...")
			default:
				fmt.Println("Waiting for data..")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}
	quit <- true
}
