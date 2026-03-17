package main

func main() {
	// simple switch

	//no need of break statement in go
	//it automatically break after executing one case

	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")
	// }

	// type switch
	// whoAmI := func(i interface{}) {
	// 	switch i.(type) { // type switch ( can only be inside switch statement with interface type)
	// 	case int:
	// 		fmt.Println("its an integer")
	// 	case string:
	// 		fmt.Println("its a string")
	// 	case bool:
	// 		fmt.Println("its a boolean")
	// 	default:
	// 		fmt.Println("other")
	// 	}
	// }

	// whoAmI(55)

	//for normal type assertion 

	// var a any
	// a = "hello"
	// str, ok := a.(string) // type assertion ( can be used anywhere)

	//if not used of 'ok' then it will panic if type assertion fails
	// str := a.(string) // panic if a is not a string
	
	// if ok {
	// 	fmt.Println("a is a string:", str)
	// } else {
	// 	fmt.Println("a is not a string")
	// }

}
