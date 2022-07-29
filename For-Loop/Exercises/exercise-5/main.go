// Problem Statement: Using a for loop print out all the years from your birthday to the current year.

// Use a variant of for loop where the post statement is moved inside the for block of code.

package main

import "fmt"

func main() {
	birthYear, currentYear := 1998, 2022

	for i := birthYear; i <= currentYear; {
		fmt.Println(i)
		i++
	}
}
