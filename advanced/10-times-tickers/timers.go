package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Create a timer for 2 seconds
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Timer started...")

	// 2. This blocks until the timer's channel 'C' receives a value
	<-timer.C
	fmt.Println("Timer expired!")

	// 3. You can also stop a timer before it expires
	stopTimer := time.NewTimer(time.Second)
	go func() {
		<-stopTimer.C
		fmt.Println("This won't print because we stop it")
	}()
	
	stopTimer.Stop() 
	fmt.Println("Timer stopped manually.")
	time.Sleep(2 * time.Second)
}