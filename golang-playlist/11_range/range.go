package main

import "fmt"

// iterating over data structures
func main() {
	nums := []int{6, 7, 8}

	for i:=0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i, num := range nums {
		fmt.Println(num, i)
	}

	//if does not want index (_ tells compiler that we are ignoring this value, so will not give error of unused variable)
	for _, num := range nums {
		fmt.Println(num)
	}

	for  num := range nums {
		fmt.Println(num)
	}

	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// 	fmt.Println(k, v)
	// }

	// if use single, then it gives only keys

	// for k := range m {
	// 	fmt.Println(k)
	// }

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
