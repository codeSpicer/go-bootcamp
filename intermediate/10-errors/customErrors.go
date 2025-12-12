package main

import (
	"errors"
	"fmt"
)

// --- 1. Custom Error Definition ---

// CustomError is a structure used to convey specific, actionable information
// beyond a simple string.
type CustomError struct {
	Operation string
	Reason    string
	Code      int
}

// Implement the built-in 'error' interface for CustomError.
// Any value of type *CustomError can now be assigned to a variable of type error.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Operation '%s' failed (Code %d): %s", e.Operation, e.Code, e.Reason)
}

// --- 2. Sentinel Error Definition (For Comparison) ---

// Sentinel errors are defined using 'var' and 'errors.New'.
// They are used to signal specific, expected failure modes throughout the code.
var ErrInvalidInput = errors.New("invalid input parameters provided")

// --- 3. Example Functions ---

// LowLevelFunction returns a sentinel error.
func LowLevelFunction(value int) error {
	if value < 0 {
		return ErrInvalidInput
	}
	// In a real scenario, this might return a network or database error
	return errors.New("database connection failed")
}

// MidLevelFunction calls the low-level function and wraps the error with context.
func MidLevelFunction(id int) error {
	err := LowLevelFunction(id)
	if err != nil {
		// Use %w to WRAP the underlying error ('err').
		// This creates an error chain (App Error -> DB Error)
		return fmt.Errorf("retrieval for ID %d failed: %w", id, err)
	}
	return nil
}

// HighLevelFunction demonstrates calling the mid-level function and handling different error types.
func HighLevelFunction(name string) error {
	// Call a function that returns a custom error struct
	if name == "" {
		return &CustomError{
			Operation: "NameValidation",
			Reason:    "Empty name string",
			Code:      1001,
		}
	}
	// Simulate success for now
	return nil
}

func customErrors() {
	fmt.Println("--- Custom Error Struct Example ---")
	if err := HighLevelFunction(""); err != nil {
		fmt.Println("Returned Error:", err)

		// errors.As: Check if the error (or any error in its chain)
		// is of a specific type (*CustomError).
		var customErr *CustomError
		if errors.As(err, &customErr) {
			fmt.Printf("HANDLED: Custom Error Type Found! Details: %s (Code: %d)\n", customErr.Reason, customErr.Code)
			// At this point, 'customErr' is a non-nil pointer, allowing access
			// to the struct's fields.
		}
	}

	fmt.Println("\n--- Wrapped Error Chain Example ---")
	appErr := MidLevelFunction(-5) // This will trigger the ErrInvalidInput sentinel

	if appErr != nil {
		fmt.Println("Full Error Chain:", appErr)

		// ----------------------------------------------------
		// 1. Unwrapping the Error (errors.Unwrap)
		// ----------------------------------------------------

		fmt.Println("\n-- 1. Unwrapping --")
		current := appErr
		for current != nil {
			fmt.Printf("Unwrapped Layer: %s\n", current.Error())
			// errors.Unwrap returns the next error in the chain, or nil if none is wrapped.
			current = errors.Unwrap(current)
		}

		// ----------------------------------------------------
		// 2. Checking the Error Chain (errors.Is)
		// ----------------------------------------------------

		fmt.Println("\n-- 2. Checking with errors.Is --")

		// errors.Is: Checks if the entire error chain contains the target sentinel error (ErrInvalidInput).
		if errors.Is(appErr, ErrInvalidInput) {
			fmt.Println("HANDLED: Found Sentinel Error (ErrInvalidInput) deep in the chain.")
			// We can decide to fix the input or log a specific warning here.
		} else {
			fmt.Println("FATAL: Unknown error type found.")
		}
	}
}
