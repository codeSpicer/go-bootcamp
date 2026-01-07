package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work.
type Task struct {
	ID    int
	Value int
}

// Result represents the outcome of a processed task.
type Result struct {
	WorkerID int
	TaskID   int
	Output   int
}

// worker is the function run by multiple goroutines concurrently.
func worker(id int, jobs <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	// 1. Ensure that when the worker finishes, it reports to the WaitGroup.
	defer wg.Done()

	for task := range jobs {
		fmt.Printf("[Worker %d] Processing Task %d...\n", id, task.ID)

		// Simulate a time-consuming operation.
		time.Sleep(500 * time.Millisecond)

		// 2. Send the result to the results channel.
		results <- Result{
			WorkerID: id,
			TaskID:   task.ID,
			Output:   task.Value * 10,
		}
	}
	fmt.Printf("[Worker %d] No more jobs. Shutting down.\n", id)
}

func main() {

	// BasicExample()

	const numWorkers = 3
	const numTasks = 6

	// Channels for communication.
	jobs := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// WaitGroup to track active workers.
	var wg sync.WaitGroup

	// --- STEP 1: Start the Workers ---
	fmt.Println("--- Starting Workers ---")
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Tell WaitGroup we are starting one worker.
		go worker(i, jobs, results, &wg)
	}

	// --- STEP 2: Send Tasks (Producer) ---
	go func() {
		for i := 1; i <= numTasks; i++ {
			jobs <- Task{ID: i, Value: i}
			time.Sleep(100 * time.Millisecond)
		}
		// Closing the jobs channel tells workers that no more tasks are coming.
		// Workers will exit their 'range jobs' loop after processing current items.
		close(jobs)
		fmt.Println("--- All Tasks Sent. Jobs Channel Closed. ---")
	}()

	// --- STEP 3: The Coordinator (The WaitGroup + Close Pattern) ---
	// We need a separate goroutine to wait for the workers to finish.
	// If we called wg.Wait() in the main thread, it would block the main
	// thread from collecting results, potentially leading to a deadlock.
	go func() {
		wg.Wait()      // Block here until all workers call wg.Done().
		close(results) // Once all workers are done, it is safe to close results.
		fmt.Println("--- All Workers Finished. Results Channel Closed. ---")
	}()

	// --- STEP 4: Collect Results (Consumer) ---
	// The main goroutine loops over results until the results channel is closed.
	fmt.Println("--- Collecting Results ---")
	for res := range results {
		fmt.Printf("[Main] Result: Task %d processed by Worker %d. Output: %d\n",
			res.TaskID, res.WorkerID, res.Output)
	}

	fmt.Println("All work completed successfully.")
}
