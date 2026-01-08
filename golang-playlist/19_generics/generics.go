package main

import "fmt"

//T any -> accept any type
//T interface{} -> accept any type ( same as any)
// T int|string -> only these two types allowed
// V string -> only string allowed

// comparable -> types that can be compared (==, !=)
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// we have to create separate functions for each type

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// LIFO
// create a generic stack data structure (using struct )
type stack[T any] struct {
	elements []T
}

func main() {
	// myStack := stack[string]{
	// 	elements: []string{"golang"},
	// }

	// fmt.Println(myStack)

	// nums := []int{1, 2, 3}
	// names := []string{"golang", "typescript"}
	vals := []bool{true, false, true}
	// printStringSlice(names)
	printSlice(vals, "john")
}
