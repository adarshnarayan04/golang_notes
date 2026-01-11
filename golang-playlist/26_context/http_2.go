package main

import (
	"context"
	"fmt"
	"time"
)

// simulateWork simulates a long-running operation that can be canceled via context.
func simulateWork(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		// Simulate work completed after 3 seconds.
		fmt.Println("work completed")
		return nil
	case <-ctx.Done():
		// Context canceled or deadline exceeded.
		fmt.Println("work canceled:", ctx.Err()) 

		// No, you cannot pass a custom error to ctx.Err().
		// ctx.Err() is set internally by the context package and only returns standard errors 
		// like context.Canceled or context.DeadlineExceeded when the context is canceled or times out.
		return ctx.Err()
	}
}

func main() {

	// Example: Call a function with context and timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := simulateWork(ctx)
	if err != nil {
		// err is the error interface value itself.
		// err.Error() is a string containing the error message.

		// Printing err (e.g., fmt.Println(err)) will call err.Error() automatically.
		// Use err.Error() if you specifically need the error message as a string (e.g., for logging, string concatenation, or custom formatting).
		// Otherwise, you can usually just use err directly in print statements.
		fmt.Println("simulateWork error:", err)
		fmt.Println(err.Error()) //always give the string message of error
	} else {
		fmt.Println("simulateWork finished successfully")
	}
}
