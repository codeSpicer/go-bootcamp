package main

import (
	"fmt"
	"time"
)

func ticker() {
	// Create a ticker that "ticks" every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	// done := make(chan bool)

	go func() {
		// Use a range loop to "wait" for every tick
		for t := range ticker.C {
			fmt.Println("Tick at", t.Format("15:04:05.000"))
		}
	}()

	// Let it run for 2 seconds
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop() // Stops the ticker
	fmt.Println("Ticker stopped")
}
