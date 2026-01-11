package main

import "fmt"

/*
	DETAILED NOTES: NAMED RETURN VALUES
	===================================
	1. What are they?
	   - Go allows you to name the return values in the function signature.
	   - These names are treated as variables defined at the top of the function.
	   - They are automatically initialized to their zero values (0, "", nil, etc.).

	2. When do we use them?
	   - Documentation: To clarify the meaning of return values, especially when multiple returns share the same type.
	     (e.g., `func GetLocation() (lat, long float64)` is clearer than `(float64, float64)`).
	   - Zero-Value Convenience: To avoid declaring variables manually before returning.
	   - Defer & Error Handling: Crucial when you need to capture or modify a return value (like an error)
	     inside a `defer` block (which runs after the return statement prepares the value but before the function exits).

	3. Naked Return:
	   - A `return` statement without arguments is called a "naked" return.
	   - It returns the current values of the named return variables.
	   - Use Rule: Use naked returns only in short functions. In long functions, they hurt readability.

	4. The Blank Identifier (_) in Named Returns:
	   - You can use `_` in the Return definition, but it prevents the use of naked returns because `_` is not a variable you can assign to.
*/

// Example 1: Basic Naked Return & Documentation
// Naming (lat, long) helps the caller understand what the float64s represent.
func getCoords() (lat, long float64) {
	lat = 37.7749
	long = -122.4194
	return // Naked return: returns current lat, long
}

// Example 2: Modifying Return Values via Defer
// A named return variable is accessible to deferred functions.
// This allows a deferred function to modify the return value *after* return is called.
func getCountWithDefer() (count int) {
	count = 10
	defer func() {
		count = count + 5 // Modifies the named return variable
		fmt.Println("Log inside defer: count increased to", count)
	}()
	return // Returns 10 + 5 = 15
}

// Example 3: Shadowing Pitfall
// A common mistake: using := (short declaration) which creates a NEW variable instead of assigning to the named return.
func shadowingRisk() (value int) {
	value = 10
	if true {
		value := 20                 // New inner variable named 'value', shadows the return variable
		fmt.Println("Inner:", value) // Prints 20
	}
	return // Returns 10 (the outer named variable), NOT 20
}

// Example 4: Named Return with Blank Identifier (_)
// If you use `_` as a name in the return signature, you strictly CANNOT use naked returns.
// You must explicitly return values.

//_ , we use this to tell , we are not going to need this variable ( no need to assign any name to it)
// but the data is still returned.
func mixedNamedReturn() (total int, _ string) {
	total = 100
	// return // COMPILATION ERROR: not enough arguments to return
	return total, "metadata-ignored" // Explain returns are required
}

// Example 5: Explicit Return (Avoiding Naked Return)
// Even with named return values, you are essentially never forced to use naked returns.
// You might choose explicit returns for clarity or to return a specific value different from the named variable.
func explicitReturnOverride(x int) (result int) {
	result = x * 2

	// Scenario: Logic dictates we should return a hardcoded value, ignoring 'result'.
	if result > 100 {
		return 0 // Explicitly returning 0, ignoring whatever is in 'result'
		//it actually set the result to 0 and return that.
	}

	// Scenario: Standard return. Using `return result` is clearer than just `return`.
	return result
}

// Example 6: Explicit Return + Defer Modification
// This demonstrates that even if you use an explicit return (not naked),
// the deferred function still operates on the named return variable
// AFTER the return statement assigns the value.
func explicitReturnWithDefer() (val int) {
	defer func() {
		// This runs AFTER 'return 5' sets val = 5
		val++ // Modifies 'val' to 6
		fmt.Println("Defer grew val to:", val)
	}()

	return 5 // Explicit return sets val to 5
}

func main() {
	fmt.Println("--- Example 1: Basic ---")
	lat, long := getCoords()
	fmt.Printf("Lat: %.4f, Long: %.4f\n", lat, long)

	fmt.Println("\n--- Example 2: Defer ---")
	c := getCountWithDefer()
	fmt.Println("Final returned count:", c)

	fmt.Println("\n--- Example 3: Shadowing ---")
	v := shadowingRisk()
	fmt.Println("Shadowing result:", v)

	fmt.Println("\n--- Example 4: Mixed with _ ---")
	t, s := mixedNamedReturn()
	fmt.Println("Mixed result:", t, s)

	fmt.Println("\n--- Example 5: Explicit Return Override ---")
	fmt.Println("Input 10 ->", explicitReturnOverride(10))   // Returns 20 (result)
	fmt.Println("Input 60 ->", explicitReturnOverride(60))   // Returns 0 (explicit override)

	fmt.Println("\n--- Example 6: Explicit Return + Defer ---")
	resDefer := explicitReturnWithDefer()
	fmt.Println("Final result:", resDefer)
}
