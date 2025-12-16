package main

import (
	"fmt"
	"time"
)

func main() {
	// --- Buffered Channel Syntax ---
	// Syntax: variable := make(chan Type, capacity)
	// Capacity (N > 0) defines the number of values that can be sent before 
	// the channel blocks, even if no receiver is ready.
	
	// Create a buffered integer channel with a capacity of 2.
	ch := make(chan int, 2)
	fmt.Printf("Channel initialized with capacity: %d\n", cap(ch))
	
	// --- 1. Demonstration: Blocking on Send if Buffer is Full ---
	fmt.Println("\n--- 1. Send Blocking Demo (Full Buffer) ---")

	// Sending 1 and 2 works immediately because the buffer has capacity 2.
	ch <- 1
	fmt.Printf("Sent 1. Current length: %d\n", len(ch))
	ch <- 2
	fmt.Printf("Sent 2. Current length: %d\n", len(ch))

	// Start a goroutine that will relieve the pressure by receiving one value later.
	go func() {
		// Wait before receiving. This allows the main routine to block first.
		time.Sleep(2 * time.Second)
		
		// Receiving frees up one slot in the buffer.
		val := <-ch
		fmt.Printf("\n[Goroutine] Received %d, buffer now has space.\n", val)
	}()

	fmt.Println("Main routine now attempting to send 3 (will block)...")
	// Sending the third value will BLOCK the main goroutine because the buffer (cap 2) is full.
	ch <- 3 
	fmt.Println("Main routine resumed (Goroutine received value, freeing buffer slot).")

	// Now that the blocking send is complete, we can receive the remaining values.
	fmt.Printf("Received %d\n", <-ch) // Receives 2
	fmt.Printf("Received %d\n", <-ch) // Receives 3
	
	// --- 2. Demonstration: Blocking on Receive if Buffer is Empty ---
	fmt.Println("\n--- 2. Receive Blocking Demo (Empty Buffer) ---")
	
	// Since we've received all values, the channel is now empty.
	// We run a goroutine to send two values after a short delay.
	go func() {
		fmt.Println("[Goroutine] Sending 1 after 1 second...")
		time.Sleep(1 * time.Second)
		ch <- 1
		fmt.Println("[Goroutine] Sent 1.")
		ch <- 2
		fmt.Println("[Goroutine] Sent 2.")
	}()

	// The main goroutine attempts to receive immediately. The channel is empty,
	// so the main routine BLOCKS until the goroutine sends the first value (after 1s).
	fmt.Println("Main routine blocking on first receive...")
	fmt.Printf("Received %d\n", <-ch) 
	
	// The channel now contains 2 (sent by the goroutine). This second receive 
	// will complete immediately without blocking.
	fmt.Printf("Received %d\n", <-ch) 

	// Final small sleep to ensure all print statements from goroutines complete.
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Program finished.")
}