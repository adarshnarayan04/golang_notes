package main
import "fmt"

//Arrays are value types.
//b = a copies the entire array

//when pass to functon, it creates a copy of the array
// numbered sequence of specific length
func main() {
	// zeroed values
	// int -> 0, string -> "", bool -> false
	// var nums [4]int

	// nums[0] = 1
	// fmt.Println(nums[0])
	// fmt.Println(nums)
	// array length
	// fmt.Println(len(nums))

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)

	// var name [3]string
	// name[0] = "golang"
	// fmt.Println(name)

	// to declare it in single line
	nums := [3]int{1, 2, 3}
	fmt.Println(nums)

	// 2d arrays
	// nums := [2][2]int{{3, 4}, {5, 6}}
	// fmt.Println(nums)

	//var nums [2]; -> have to specify the type 
	//fmt.Println(nums)




	// - fixed size, that is predictable
	// - Memory optimazation
	// - Contant time access

}
