package main

import "fmt"

func main() {
	var hello =10;
	fmt.Println(hello)
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// Declaration + init (type inferred)
	//Works at package or function scope.
	var a = 10

	// Short declaration (function scope only)
	//  only allowed inside functions.
	b := 20

	// Declaration then assignment
	var c int
	c = 30

	// Explicit type to control the type
	var d float32 = 3.0

	fmt.Println(a, b, c, d)
}
