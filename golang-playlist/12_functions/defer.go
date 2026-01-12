package main

import (
	"fmt"
	"os"
)

// --- NOTES: defer in Go ---
//
// 1. What is defer?
//    - The defer statement postpones the execution of a function until the surrounding function returns.
//    - Commonly used for cleanup: closing files, unlocking mutexes, etc.
//
// 2. How does it work?
//    - Deferred calls are executed in LIFO (Last-In, First-Out) order.
//    - Arguments to deferred functions are evaluated immediately, but the function call is delayed.
//
// 3. Typical use cases:
//    - Resource cleanup (file.Close, unlock, etc.)
//    - Logging exit points
//    - Tracing function execution

// --- Example 1: Basic defer ---
func basicDefer() {
	fmt.Println("Start")
	defer fmt.Println("Deferred: runs last")
	fmt.Println("End")
}

// --- Example 2: Multiple defers (LIFO order) ---
func multipleDefers() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	fmt.Println("Function body")
	// Output order: Function body, Second defer, First defer
}

// --- Example 3: Defer for resource cleanup ---
func fileDefer() {
	f, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close() // Ensures file is closed when function exits
	fmt.Fprintln(f, "Hello, defer!")
}

// --- Example 4: Arguments evaluated immediately ---
//only argumnets are evaluated immediately, not the function itself
func argDefer() {
	x := 10
	defer fmt.Println("Deferred value of x:", x)
	x = 20
	fmt.Println("Current value of x:", x)
	return
	// Output: Deferred value of x: 10
}

func deferNamedReturn() (result int) {
	result = 10
	defer func(result int) {
		result++ // This modifies the local copy of result, not the named return value!
		fmt.Println("Inside defer, modified result:", result)
	}(result) //it passes the current value of result (10) to the deferred function ( as arguments are evaluated immediately)
	result = 12
	return // result will be 12, NOT 13
}

// --- Example 5a: Defer modifying named return value using pointer ---
func deferNamedReturnPtr() (result int) {
	result = 10
	defer func(r *int) {
		*r++ // Modifies the named return value via pointer
		fmt.Println("Inside defer (pointer), modified result:", *r)
	}(&result) //useful when function code is not written here ( so then it can not access the named return variable directly), so have to pass
	result = 12
	return // result will be 13 (12 + 1 from defer)
}

// --- Example 5b: Defer modifying named return value using closure ---
func deferNamedReturnClosure() (result int) {
	result = 10
	defer func() {
		result++ // Modifies the named return value directly
		fmt.Println("Inside defer (closure), modified result:", result)
	}()
	result = 12
	return // result will be 13 (12 + 1 from defer)
}

// --- NOTES: Function chaining in defer ---
// You can chain function calls in a defer statement.
// All arguments are evaluated immediately, so the result of the chain is fixed at the time of defer.

// --- Example 6: Function chaining in defer ---
func chainedDefer() {
	fmt.Println("Start chainedDefer")
	res := 10
	defer fmt.Println("Deferred:", double(increment(res))) // increment(5) is called immediately, then double()
	//as double(inrement(res)) is argument to fmt.Println, so both increment and double are called immediately
	res = 12
	fmt.Println("End chainedDefer")
}

func increment(x int) int {
	fmt.Println("increment called with", x)
	return x + 1
}

func double(x int) int {
	fmt.Println("double called with", x)
	return x * 2
}

// --- Example 7: Method chaining in defer ---
type Chain struct {
	val int
}

func (c Chain) Inc() Chain {
	fmt.Println("Inc called with", c.val)
	return Chain{val: c.val + 1}
}

func (c Chain) Double() Chain {
	fmt.Println("Double called with", c.val)
	return Chain{val: c.val * 2}
}

func (c Chain) Print() {
	fmt.Println("Chain value:", c.val)
}

func chainedMethodDefer() {
	fmt.Println("Start chainedMethodDefer")
	c := Chain{val: 5}
	defer c.Inc().Double().Print() // Inc and Double are called immediately, Print is deferred
	//as inc and double are areguments to Print method, so both are called immediately
	// as arguments are evaluated immediately
	c.val = 10
	fmt.Println("End chainedMethodDefer")
}

func main() {
	basicDefer()
	multipleDefers()
	fileDefer()
	argDefer()
	fmt.Println("deferNamedReturn result:", deferNamedReturn())
	fmt.Println("deferNamedReturnPtr result:", deferNamedReturnPtr())
	fmt.Println("deferNamedReturnClosure result:", deferNamedReturnClosure())
	chainedDefer()
	chainedMethodDefer()
}
