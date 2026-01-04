package main

import (
	"fmt"
	"time"
)

func main() {
	// Pattern 1: Handling a stream of data with a timeout and close detection
	fmt.Println("--- Pattern 1: Loop with Select and Timeout ---")
	pattern1_LoopAndTimeout()

	// Pattern 2: Multi-channel selection (The first one to respond wins)
	fmt.Println("\n--- Pattern 2: Multi-Channel Competition ---")
	pattern2_MultiChannel()

	// Pattern 3: Non-blocking channel operations using 'default'
	fmt.Println("\n--- Pattern 3: Non-blocking Select ---")
	pattern3_NonBlocking()
}

// pattern1_LoopAndTimeout demonstrates how to keep a receiver alive 
// until a channel is closed, while also preventing hangs via a timeout.
func pattern1_LoopAndTimeout() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i
		}
		// Closing is essential for the receiver to know when to stop.
		close(ch) 
	}()

	for {
		select {
		case msg, ok := <-ch:
			// The 'ok' variable is false if the channel is empty AND closed.
			if !ok {
				fmt.Println("[Receiver] Channel closed. Exiting loop.")
				return
			}
			fmt.Println("[Receiver] Received:", msg)

		case <-time.After(2 * time.Second):
			// This case triggers if no other case is ready for 2 seconds.
			fmt.Println("[Receiver] Timeout: No data received for too long.")
			return
		}
	}
}

// pattern2_MultiChannel shows how select waits for multiple sources.
// If multiple channels are ready at the same time, Go picks one RANDOMLY.
func pattern2_MultiChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from Channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from Channel 2"
	}()

	// We only want to handle the FIRST message that arrives.
	select {
	case msg1 := <-ch1:
		fmt.Println("[Main] Winner:", msg1)
	case msg2 := <-ch2:
		fmt.Println("[Main] Winner:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("[Main] Error: Both sources timed out.")
	}
}

// pattern3_NonBlocking demonstrates the 'default' case.
// If no channel has a value ready, 'default' executes immediately 
// rather than blocking and waiting.
func pattern3_NonBlocking() {
	ch := make(chan int)

	// We don't start a goroutine here, so the channel will be empty.

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		// Because there's no sender, the 'case' above is not ready.
		// 'default' prevents the program from deadlocking.
		fmt.Println("[Main] No data available right now. Moving on...")
	}
}