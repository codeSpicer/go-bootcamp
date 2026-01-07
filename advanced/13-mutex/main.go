package main

import (
	"fmt"
	"sync"
)

// SafeCounter wraps an integer and a Mutex to ensure thread-safe access.
type SafeCounter struct {
	value int
	// RWMutex allows multiple readers OR one single writer.
	mu sync.Mutex
}

// Increment safely increases the counter by 1.
func (c *SafeCounter) Increment() {
	// Lock() ensures only one goroutine can enter this section at a time.
	c.mu.Lock()
	// Using defer for Unlock ensures the mutex is released even if the function panics.
	defer c.mu.Unlock()

	c.value++
}

// Value safely returns the current counter value.
func (c *SafeCounter) Value() int {
	// RLock (Read Lock) allows other readers to access the value simultaneously,
	// but blocks any writers. This is more efficient for "Read-Heavy" apps.
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// 1. Thread-Safe instance
	safeNum := SafeCounter{value: 0}

	// 2. Unprotected instance (This will suffer from a Data Race)
	unsafeNum := 0

	var wg sync.WaitGroup
	numGoroutines := 10
	incrementsPerRoutine := 1000

	fmt.Printf("Starting %d goroutines, each incrementing %d times...\n", numGoroutines, incrementsPerRoutine)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerRoutine; j++ {
				// Protected: Only one goroutine can write at a time.
				safeNum.Increment()

				// Unprotected: Multiple goroutines try to read/write this memory
				// at the same time. Some updates will be lost!
				unsafeNum++
			}
		}()
	}

	// Wait for all work to complete.
	wg.Wait()

	fmt.Println("\n--- Results ---")
	fmt.Printf("Safe Counter Result:   %d (Expected: %d)\n", safeNum.Value(), numGoroutines*incrementsPerRoutine)
	fmt.Printf("Unsafe Counter Result: %d (Likely incorrect due to Data Race)\n", unsafeNum)

	if unsafeNum != numGoroutines*incrementsPerRoutine {
		fmt.Println("\nNotice: The unsafe counter failed because multiple goroutines")
		fmt.Println("overwrote each other's work simultaneously.")
	}
}
