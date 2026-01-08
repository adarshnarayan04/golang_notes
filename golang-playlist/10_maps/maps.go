package main

// maps -> hash, object, dict
func main() {

	//wrong way
	// var m map[string]int
	// m["apple"] = 10
	// panic: assignment to entry in nil map
	// as m is nil map , does not point to any underlying data structure in heap
	// to make it point to underlying data structure we have to use make
	
	//right way
	// var m = make(map[string]int)
	//or 
	//var m map[string]int
	// m = make(map[string]int) -> it make m point to underlying data structure in heap
	// or var m= make(map[string]int)

	// as make allocates memory in heap and returns the reference to it
	// var m= make(map[string]int) , assigns the reference to m
	// m["apple"] = 10
	// fmt.Println(m["apple"])

	// creating map

	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])
	// IMP: if key does not exists in the map then it returns zero value

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))

	// delete(m, "price")
	// clear(m)

	// fmt.Println(m)
	// fmt.Println(m)

	//without make
	// m := map[string]int{} 
	// var m= map[string]int{}
	// as {} initializes the map a, so m points to underlying data structure in heap ,as it is initialized


	// m := map[string]int{"price": 40, "phones": 3}

	// v, ok := m["phones"]
	// ok -> boolean value , tells whether the key is present in map or not , value associated with key
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// m1 := map[string]int{"price": 40, "phones": 3}
	// m2 := map[string]int{"price": 40, "phones": 8}
	// fmt.Println(maps.Equal(m1, m2))

}
