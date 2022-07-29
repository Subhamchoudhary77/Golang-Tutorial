package main

import "fmt"

func main() {

	// append in same slice
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5)
	fmt.Println(numbers)

	// append in different slice
	count := []int{6, 7, 8}
	numbers = append(numbers, count...)
	fmt.Println(numbers)

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices.
	src := []int{11, 12, 13, 14, 15}
	dst := make([]int, len(src))
	counter := copy(dst, src)
	fmt.Println(dst)
	fmt.Println(src, dst, counter)

	//** Slice's Length and Capacity **//

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 1, Capacity: 1

	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 5, Capacity: 8

}
