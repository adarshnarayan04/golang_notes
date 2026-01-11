package main

import (
	"fmt"
	"net/http"
	"time"
)

// hello is an HTTP handler function.
// It simulates a long-running operation and demonstrates how to use context.Context
// to handle request cancellation.
func hello(w http.ResponseWriter, req *http.Request) {

	// A context.Context is created for each request by the net/http package.
	// You can get it from the request using req.Context().
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Simulate some work by waiting for 10 seconds before replying.
	// While waiting, check if the context is "done" (i.e., if the client canceled the request).
	select {
	case <-time.After(10 * time.Second):
		// If 10 seconds pass without cancellation, send a normal response.
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done()://ctx.Done() is a receive-only channel provided by the context implementation(package). It is closed by the context itself when it is canceled or times out.
		//Wait for a few seconds before sending a reply to the client. This could simulate some work the server is doing. While working, keep an eye on the context’s Done() channel for a signal that we should cancel the work and return as soon as possible.
		// ctx.Done() returns a channel that's closed when the context is canceled
		// (for example, if the client disconnects or cancels the request).
		// When this happens, we should stop processing and return as soon as possible.

		// ctx.Err() returns an error explaining why the context was canceled.
		// This error is usually context.Canceled or context.DeadlineExceeded.
		err := ctx.Err()
		fmt.Println("server:", err) // Log the error on the server side.

		// Send an HTTP 500 error to the client with the error message.
		// http.Error writes a response with the specified status code and error message.
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	fmt.Println("Server is listening on port 8090...")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}