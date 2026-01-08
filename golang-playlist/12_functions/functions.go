package main

// func add(a int , b int) int {
// 	return a + b
// }

//shorthand way (if parameters are of same type)
// func add(a, b int) int {
// 	return a + b
// }

// can have multiple return values
// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true
// }

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {

	//Anonymous Function Assigned to a Variable

	// fn := func(a int) int {
	// 	return 2
	// }

	//IIFE
	// func() {
    // fmt.Println("Hello from anonymous function")
	// }()

	fn := processIt()
	fn(6)

	// lang1, lang2, _ := getLanguages()
	// fmt.Println(lang1, lang2)
	// result := add(3, 5)
	// fmt.Println(result)
}
