package main

import (
	"context"
	"fmt"
	"time"
)

/*
	DETAILED NOTES: CONTEXT IN GO
	=============================
	1. What is Context?
	   - The `context` package provides a standardized way to handle:
	     a) Cancellation signals.
	     b) Deadlines & Timeouts.
	     c) Request-scoped values (like UserID, Auth Tokens).
	   - It is designed to be passed down the call stack (from main -> handler -> db -> etc.).

	2. Key Interfaces & Functions:
	   - context.Background(): The root context (empty, never handled). Used in main/init.
	   - context.TODO(): Placeholder when you don't know which context to use yet.
	   - context.WithCancel(parent): Returns a copy of parent that closes its Done channel when `cancel()` is called.
	   - context.WithTimeout(parent, duration): Cancels automatically after duration.
	   - context.WithDeadline(parent, time): Cancels automatically at a specific time.
	   - context.WithValue(parent, key, value): key-value store for request-scoped data.

	3. Real-Life Use Cases:
	   - HTTP Servers: Each request gets a context. If the client disconnects, the context is canceled, avoiding wasted work (long DB queries).
	   - Database Calls: Set timeouts to prevent hanging queries.
	   - Microservices: Propagate Trace IDs (OpenTelemetry) across service boundaries.
	   - Graceful Shutdown: Cancel all running background workers when the server shuts down.

	4. Best Practices:
	   - Always pass `ctx` as the FIRST argument to a function.
	   - Never store Context inside a struct type; pass it explicitly.
	   - Use context values only for request-scoped data (not for optional parameters).
*/

// Example 1: WithCancel
// Simulates a long-running worker that stops when we tell it to.
func exampleWithCancel() {
	fmt.Println("--- Example 1: WithCancel ---")
	// Create a context that can be manually canceled
	ctx, cancel := context.WithCancel(context.Background())

	//ctx.Done() is only recieve-only channel, we can only listen to it. ( as it is create by context package)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // Listens for cancellation signal
				fmt.Println("Worker: I received a cancel signal. Stopping.")
				return
			default:
				fmt.Println("Worker: Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Main: Cancelling the worker now...")
	cancel() // Sends signal to ctx.Done()
	//now we can cancel the function using ctx.Done() channel provided by context package.
	time.Sleep(1 * time.Second) // Wait to see worker stop
}

// Example 2: WithTimeout
// Simulates an API call that must finish within a specific time.
func exampleWithTimeout() {
	fmt.Println("\n--- Example 2: WithTimeout ---")
	// Context automatically cancels after 1 second
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Good practice to call cancel context to release resources, even if timeout handles it

	resultChan := make(chan string)

	go func() {
		// Simulate slow operation (2 seconds)
		time.Sleep(2 * time.Second)
		resultChan <- "Operation Complete"
	}()

	select {//select the channel which is ready first
		// so if we get the resilt from function first we print that
		// otherwise if context is done (timeout) , then 
	case res := <-resultChan:
		fmt.Println("Success:", res)
	case <-ctx.Done():
		fmt.Println("Error: Operation timed out!", ctx.Err())
	}
}

// Alternative Example 2: WithTimeout without using a channel
func exampleWithTimeoutNoChannel() {
	fmt.Println("\n--- Example 2 (No Channel): WithTimeout ---")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		// Simulate slow operation (2 seconds), but check context in loop
		for i := 0; i < 4; i++ {
			select {
			case <-ctx.Done():
				// Context canceled, exit early
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
		// Work completed
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("Success: Operation Complete")
	case <-ctx.Done():
		fmt.Println("Error: Operation timed out!", ctx.Err())
	}
}

// Example 3: WithDeadline
// Similar to Timeout, but uses a specific wall-clock time.
func exampleWithDeadline() {
	fmt.Println("\n--- Example 3: WithDeadline ---")
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-time.After(2 * time.Second): // Simulate task taking 2s
		fmt.Println("Task finished")
	case <-ctx.Done():
		fmt.Println("Error: Deadline exceeded!", ctx.Err())
	}
}

// Example 4: WithValue
// Passing request-scoped data like a User ID.
func exampleWithValue() {
	fmt.Println("\n--- Example 4: WithValue ---")
	// Key should practically be a custom type to avoid collisions, using string here for simplicity
	type keyType string
	const userKey keyType = "userID"

	ctx := context.WithValue(context.Background(), userKey, 12345)

	processRequest(ctx, userKey)
}

// Example 5: Passing Multiple Values (Chaining)
// context.WithValue only adds ONE key-value pair. To add multiple, you wrap the context repeatedly.
// Every time you call WithValue, it returns a NEW context that wraps the parent.
func exampleMultipleValues() {
	fmt.Println("\n--- Example 5: Multiple Values (Chaining) ---")
	type keyType string
	const (
		keyUserID keyType = "userID"
		keyRole   keyType = "role"
	)

	// Step 1: Create base context
	ctx := context.Background()

	// Step 2: Add first value
	ctx = context.WithValue(ctx, keyUserID, 9876)

	// Step 3: Add second value (wraps the previous context)
	// You can chain this indefinitely, though too many layers can impact lookup performance.
	ctx = context.WithValue(ctx, keyRole, "admin")

	// Retrieve values
	valID := ctx.Value(keyUserID)
	valRole := ctx.Value(keyRole)

	fmt.Printf("User ID: %v, Role: %v\n", valID, valRole)
}

// Example 6: Passing Complex Data (Struct)
// Instead of chaining many keys (which can get messy), it's often cleaner to pass a single struct.
func exampleStructInContext() {
	fmt.Println("\n--- Example 6: Struct in Context ---")
	type keyType string
	const authKey keyType = "authData"

	type AuthData struct {
		UserID   int
		Username string
		IsAdmin  bool
	}

	data := AuthData{UserID: 555, Username: "gopher", IsAdmin: true}

	// Store the whole struct under one key
	ctx := context.WithValue(context.Background(), authKey, data)

	// Retrieve: We must type-assert the value back to AuthData to use the fields
	if val, ok := ctx.Value(authKey).(AuthData); ok {
		fmt.Printf("Authenticated User: %s (ID: %d, Admin: %t)\n", val.Username, val.UserID, val.IsAdmin)
	}
}

func processRequest(ctx context.Context, key any) {
	if val := ctx.Value(key); val != nil {
		fmt.Printf("Found value in context: %v\n", val)
	} else {
		fmt.Println("Value not found")
	}
}

func main() {
	exampleWithCancel()
	exampleWithTimeout()
	exampleWithTimeoutNoChannel() // Call the new version
	exampleWithDeadline()
	exampleWithValue()
	exampleMultipleValues()
	exampleStructInContext()
}
