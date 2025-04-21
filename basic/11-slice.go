package basic

import "fmt"

//slice is dynamic aray
//no need to give length of aaray in avance
//most used construct in go
//useful array methods

func SliceDemo() {
	fmt.Println("this is slice demo")
	var nums []int = []int{1, 2, 3, 4}
	fmt.Println(nums)
	fmt.Println(len(nums))

	/*
	*nil slice
	 */
	var nums2 []int
	fmt.Println(nums2 == nil)

	/*
	* not nill slice
	 */

	var nums3 = make([]int, 5, 10) //10 is max number of elements and this is cap

	fmt.Println(nums3) //[0 0 0 0 0]

	/*
	* adding ememts to array
	 */
	nums3 = append(nums3, 21, 1, 2, 2, 12, 1, 3, 4)
	fmt.Println(nums3)      //[0 0 0 0 0 21]
	fmt.Println(cap(nums3)) //20
	fmt.Println(len(nums3)) //13

	/*
	*making copy of slice nums3
	 */
	original := []int{1, 2, 3, 4}
	clone := make([]int, len(original)) // Create a new slice with the same length
	fmt.Println("Clone:", clone)
	copy(clone, original) // Copy elements from original to clone

	fmt.Println("Original:", original)
	fmt.Println("Clone:", clone)

	// Modifying the clone won't affect the original
	clone[0] = 100
	fmt.Println("After modification:")
	fmt.Println("Original:", original)
	fmt.Println("Clone:", clone)
}
