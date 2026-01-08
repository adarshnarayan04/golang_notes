package main
import "fmt"

//refrece data type
//go have three reference data types
//1. slices
//2. maps
//3. channels
//we use make to create these reference types , as they point to underlying data structure in heap

// slice -> dynamic
// most used construct in go
// + useful methods
// we declaire slice using [] , as it is dynamic we dont specify size

//when pass to function it pass the reference not copy of slice ( as it is reference type)
func main() {
	// uninitialized slice is nil ( as it size is not defined , so when we declare it, it does not point to any underlying array in heap)
	// var nums []int
	// var nums []int declares a slice without initializing it. That slice is nil.
	// means nums does not point to any underlying array(in heap).
	//to make it point to array(heap) we have assign it num


	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var s []int       // nil slice
    // s2 := []int{}     // empty slice
	//it create empty slice , also s2 is initialized and points to underlying array in heap
    // s3 := make([]int, 0)

	//by make -> we can create slice, map ... as they are reference types
	// make -> allocates memory in heap and returns the reference to it
	// if we size 3 , then [3]int -> it is array of size 3 not slice
	//for slice we have to use make([]int, 3)

	// var nums = make([]int, 0, 5) // length , capacity
	// capacity -> maximum numbers of elements can fit

	//when the length exceeds capacity go automatically increase the capacity
	//capacity is doubled , when this happen (like in vectors in c++)

	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)

	// nums := []int{}
	// {} -> means its initialized with empty size ( as in := we have to initialize it)

	// nums = append(nums, 1)
	// nums = append(nums, 2)

	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// // copy function ( as slices are reference types not value types so we cant do direct assignment)
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator

	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:1]) //[start:end] end is exclusive
	// fmt.Println(nums[:2]) // from start to index 2 (2 exclusive) ( default start is 0)
	// fmt.Println(nums[1:])

	// slices
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2))

	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)

	var s []int       // nil slice
	s2 := []int{}  // empty slice
	//s4 := []int   //:= both declares and initializes a variable — and []int by itself is only a type, not a value.
	// that why this give error
	s3 := make([]int, 0)

	fmt.Println(s, s2, s3)


}
