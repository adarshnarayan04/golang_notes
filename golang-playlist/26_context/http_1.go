package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// --- Approach 1: Create a context with timeout and attach it to the request ---
	// This allows you to cancel the request if it takes too long.
	// The context is attached to the http.Request and can be accessed on the server side via req.Context().
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8090/hello", nil)
	if err != nil {
		fmt.Println("failed to create request:", err)
		return
	}

	// req.Context() returns the context associated with this request.
	// On the server side, handlers can use req.Context() to detect cancellation, timeouts, or pass values.
	// If you use http.NewRequestWithContext, your custom context (with timeout/cancel) is used.
	// If you use http.NewRequest, the default context.Background() is used.

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// If the request takes longer than 2 seconds, context will cancel it and return an error.
		fmt.Println("request failed (with context):", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("response status (with context):", resp.Status)
	}

	// --- Approach 2: Use the default context (no timeout/cancel) ---
	// This request will not be canceled automatically; it will wait as long as needed.
	// The context in req2 will be context.Background().
	req2, err := http.NewRequest(http.MethodGet, "http://localhost:8090/hello", nil)
	if err != nil {
		fmt.Println("failed to create request:", err)
		return
	}

	resp2, err := client.Do(req2)
	if err != nil {
		fmt.Println("request failed (default context):", err)
		return
	}
	defer resp2.Body.Close()
	fmt.Println("response status (default context):", resp2.Status)

	// --- Approach 3: Manually cancel the request after 2 seconds using context cancel function ---
	// This demonstrates how you can cancel a request yourself (not just by timeout).
	ctx3, cancel3 := context.WithCancel(context.Background())
	req3, err := http.NewRequestWithContext(ctx3, http.MethodGet, "http://localhost:8090/hello", nil)
	if err != nil {
		fmt.Println("failed to create request:", err)
		return
	}

	done := make(chan struct{})
	go func() {
		resp3, err := client.Do(req3)
		if err != nil {
			fmt.Println("request failed (manual cancel):", err)
		} else {
			defer resp3.Body.Close()
			fmt.Println("response status (manual cancel):", resp3.Status)
		}
		close(done)
	}()

	time.Sleep(2 * time.Second) // Wait for 2 seconds before canceling
	cancel3()                   // Manually cancel the request
	<-done                      // Wait for the goroutine to finish
}
