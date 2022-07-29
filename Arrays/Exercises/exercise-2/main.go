package main

import "fmt"

func main() {
	nums := [...]int{30, -1, -6, 90, -6}

	// Problem Statement: Consider the above array, and Write the logic that counts the number of positive even numbers in the array.
	count := 0
	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			count++
		}
	}
	fmt.Printf("Number of Positive Numbers are: %d\n", count)
}
