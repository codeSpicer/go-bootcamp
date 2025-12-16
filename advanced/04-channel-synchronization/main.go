package main

import (
	"fmt"
	"time"
)

func main() {
	// --- Single Goroutine Synchronization Examples ---

	// Renamed from example1: Demonstrates using an unbuffered channel to wait
	// for a single concurrent task to finish.
	signalCompletion()

	// Renamed from example2: Demonstrates transferring a value between two goroutines.
	transferData()

	// The following sleep is crucial to ensure the asynchronous goroutine
	// in transferData() has time to execute its final print statement before
	// the main program terminates.
	time.Sleep(100 * time.Millisecond)

	// --- Synchronizing Multiple Goroutines (The Worker Pool Pattern) ---

	numGoroutines := 3
	// Use a buffered channel large enough to hold all completion signals (cap 3).
	// This prevents the workers from blocking immediately if the main routine is
	// slow to start receiving.
	done := make(chan int, numGoroutines)

	fmt.Println("\n----- Synchronizing Multiple Workers -----")

	// 1. Start all worker goroutines concurrently.
	for i := range numGoroutines {
		// Use a function literal and pass 'i' as an argument (id) to avoid the
		// closure variable bug where all goroutines might use the final value of 'i'.
		go func(id int) {
			fmt.Printf("[Worker %d] Working...\n", id)
			time.Sleep(1 * time.Second) // Simulate task duration

			// Signal Completion: Send the worker's ID back on the channel.
			done <- id
		}(i)
	}

	// 2. Block and Wait for all signals.
	// We range over the number of workers, blocking on '<-done' in each iteration.
	for range numGoroutines {
		completedId := <-done
		fmt.Printf("[Worker %d] Completed.\n", completedId)
	}
	fmt.Println("All workers synchronized and completed.")

	// --- Synchronizing Data Exchange and Closing Channels (Range Loop) ---

	fmt.Println("\n----- Synchronizing Data Exchange (Range Loop) -----")

	// Unbuffered channel for sequential data transfer.
	data := make(chan string)

	// Producer Goroutine: Sends data and closes the channel.
	go func() {
		for i := range 5 {
			data <- fmt.Sprintf("Data-%d", i) // Send formatted data
			time.Sleep(100 * time.Millisecond)
		}
		// Closing the channel is mandatory when using 'for range' on the receiver side.
		// It signals to the receiver that no more values will be sent.
		close(data)
	}()

	// Consumer (Main Routine): Uses 'for range' to receive data.
	// The loop automatically breaks (without error) when the channel is closed.
	for value := range data {
		fmt.Printf("[Consumer] Received: %s at %v\n", value, time.Now().Format("15:04:05.000"))
	}
	fmt.Println("Data exchange completed. Channel closed by sender.")

	fmt.Println("\n--- Main function execution finished ---")
}

// signalCompletion demonstrates synchronization using an empty struct channel (struct{}).
// The empty struct takes up zero memory, making it ideal for signaling simple events.
func signalCompletion() {
	fmt.Println("\n--- Single Task Synchronization (Signal Completion) ---")

	// done is an unbuffered channel used to signal completion.
	done := make(chan struct{})

	// Start a concurrent task.
	go func() {
		fmt.Println("[Task] Task started...")
		time.Sleep(time.Second) // Simulate work being done

		// Signal: Send the zero-value struct. This blocks if no receiver is ready.
		done <- struct{}{}
	}()

	// Wait: Main goroutine blocks here until a signal is received from 'done'.
	<-done
	fmt.Println("[Main] Task completed, synchronization achieved.")
}

// transferData demonstrates unbuffered channel communication (transferring a value).
func transferData() {
	fmt.Println("\n--- Single Value Communication (Transfer Data) ---")

	// ch is an unbuffered channel for integer data.
	ch := make(chan int)

	// Sender goroutine
	go func() {
		// Send: This operation BLOCKS until the receiver is ready.
		ch <- 4

		// This sleep and print happens *after* the value 4 is successfully transferred
		// and the sender is unblocked.
		time.Sleep(50 * time.Millisecond)

		// This print line relies on the main function staying alive long enough.
		fmt.Println("[Sender] Value sent and goroutine finished.")
	}()

	// Receiver: Main goroutine blocks here until the value is sent.
	value := <-ch

	// These lines execute immediately after the transfer unblocks the main routine.
	fmt.Println("[Main] Value received.")
	fmt.Println("[Main] Received value:", value)
}
