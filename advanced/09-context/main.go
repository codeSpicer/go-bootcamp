package main

import (
	"context"
	"fmt"
	"time"
)

// In Go, context is ALWAYS the first parameter of a function.
// It is usually named 'ctx'.
func performTask(ctx context.Context, taskName string) {
	for {
		select {
		case <-ctx.Done():
			// ctx.Done() is a channel that closes when:
			// 1. The cancel() function is called.
			// 2. The timeout/deadline expires.
			fmt.Printf("[%s] Stopping: %v\n", taskName, ctx.Err())
			return
		default:
			// Simulate work
			fmt.Printf("[%s] In progress...\n", taskName)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// --- 1. context.Background() and context.TODO() ---
	// Background: Use this as the root for your main function or top-level requests.
	// TODO: Use this if you aren't sure which context to use yet (a placeholder).
	rootCtx := context.Background()

	// --- 2. context.WithTimeout ---
	// Used for: API calls, Database queries, or any task with a time limit.
	fmt.Println("--- Starting Timeout Example ---")
	timeoutCtx, cancelTimeout := context.WithTimeout(rootCtx, 2*time.Second)

	// ALWAYS call cancel(). If the task finishes BEFORE 2 seconds,
	// calling cancel() releases resources associated with the context.
	defer cancelTimeout()

	// This will run for 2 seconds and then stop automatically.
	go performTask(timeoutCtx, "TimeoutTask")
	time.Sleep(2500 * time.Millisecond)

	// --- 3. context.WithCancel ---
	// Used for: Manual cancellation (e.g., a user clicks 'Cancel' or the app is shutting down).
	fmt.Println("\n--- Starting Manual Cancel Example ---")
	cancelCtx, cancelManual := context.WithCancel(rootCtx)

	go performTask(cancelCtx, "ManualTask")

	time.Sleep(1 * time.Second)
	fmt.Println("[Main] Decided to cancel manual task now...")
	cancelManual() // Manually trigger the <-ctx.Done() case
	time.Sleep(100 * time.Millisecond)

	// --- 4. context.WithValue ---
	// Used for: Request-scoped metadata (User IDs, Request IDs, Auth Tokens).
	// DANGER: Never use this for "optional arguments" or configuration.
	// It's for data that belongs to the specific execution flow.
	fmt.Println("\n--- Starting Value Example ---")

	// Keys should ideally be custom types to avoid collisions, but strings work for basics.
	type contextKey string
	const userKey contextKey = "userID"

	valueCtx := context.WithValue(rootCtx, userKey, "user_12345")

	identifyUser(valueCtx, userKey)
}

func identifyUser(ctx context.Context, key any) {
	// Value retrieval returns 'any', so you must type-assert it.
	// if userID, ok := ctx.Value(key).(string); ok {
	// 	fmt.Printf("Found User ID in context: %s\n", userID)
	// } else {
	// 	fmt.Println("No User ID found in context.")
	// }
	if ctx.Value(key) != nil {
		fmt.Printf("Found User ID in context: %s\n", ctx.Value(key))
	} else {
		fmt.Println("No User ID found in context.")
	}
}
