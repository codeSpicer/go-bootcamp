package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// sum is a goroutine function that calculates the sum of an integer slice
// and sends the result back on the provided channel 'ch'.
func sum(arr []int, ch chan int) {
	var totalSum int
	for _, v := range arr {
		totalSum += v
	}
	// Send the calculated sum back to the main goroutine via the channel.
	// This operation is blocking until the main routine receives the value.
	ch <- totalSum
}

func main() {
	// Channels are the primary mechanism used to enable communication and
	// synchronization between concurrent goroutines.
	// Syntax: variable := make(chan Type)

	// --- 1. Demonstration of Communication (Producer/Consumer Pattern) ---

	// Create an unbuffered string channel. Operations will block until
	// both a sender and a receiver are ready.
	charChannel := make(chan string)

	// Producer Goroutine: Sends characters one by one.
	go func() {
		// Iterate over a string literal, which yields runes. Convert back to string.
		for _, charRune := range "abcde" {
			// Sending blocks until the receiver is ready.
			charChannel <- string(charRune)
			// Add a short delay to better observe the staggered output
			time.Sleep(100 * time.Millisecond)
		}
		// NOTE: In a real scenario, you would close(charChannel) here,
		// but we rely on the receiver stopping after 6 loops below.
	}()

	// Consumer Goroutine: Receives and prints characters.
	go func() {
		// Attempt to receive 6 items, even though the sender only sends 5.
		// The 6th receive attempt will cause the whole program to deadlock
		// or wait indefinitely unless the main routine closes the channel or exits.
		for i := range 6 {
			// Receiving blocks until the sender sends a value.
			receivedChar := <-charChannel
			// Print the received item and the iteration index.
			fmt.Printf("[Consumer] Received: %s (Index: %d)\n", receivedChar, i)
		}
	}()

	// --- 2. Demonstration of Synchronization (Splitting Workload) ---

	// Create an unbuffered integer channel to collect partial sums.
	partialSums := make(chan int)

	// 2a. Initialize data: Create a slice of 100 random numbers.
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = rand.IntN(100) // Generate random numbers from 0 to 99
	}

	// 2b. Start two goroutines to perform parallel sum calculations.

	// Goroutine 1: Sum the first half of the array (arr[0] to arr[49])
	go sum(arr[:len(arr)/2], partialSums)

	// Goroutine 2: Sum the second half of the array (arr[50] to arr[99])
	go sum(arr[len(arr)/2:], partialSums) // Corrected slice indexing to include the 50th element

	// 2c. Receiving the results (Synchronization point).
	// The main routine BLOCKS here until the two 'sum' goroutines
	// send their results back on the 'partialSums' channel.
	x, y := <-partialSums, <-partialSums

	fmt.Println("\n--- Synchronization Results ---")
	fmt.Printf("Partial Sum 1 (x): %d\n", x)
	fmt.Printf("Partial Sum 2 (y): %d\n", y)
	fmt.Printf("Total Sum (x + y): %d\n", x+y)

	// Keep the main goroutine alive long enough for the asynchronous
	// Producer/Consumer routines (Section 1) to finish their work.
	// Without this, the program would exit immediately after calculating the sum.
	fmt.Println("\nWaiting for Producer/Consumer to finish...")
	time.Sleep(1 * time.Second)
	fmt.Println("Program finished.")
}
