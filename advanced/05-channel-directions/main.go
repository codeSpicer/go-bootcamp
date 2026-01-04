package main

import "fmt"

func main() {

	ch := make(chan int) // bi directional channel

	sendData(ch)
	receiveData(ch)

}

func sendData(ch chan<- int) {
	go func(ch chan<- int) { // send only channel
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}(ch)
}

func receiveData(ch <-chan int) { // receive only channel
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}
