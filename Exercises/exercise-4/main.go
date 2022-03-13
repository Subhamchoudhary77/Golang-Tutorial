package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// 1. Print each variable using a specific verb for its type

	fmt.Printf("Value of x is %d\n", x)
	fmt.Printf("Value of y is %f\n", y)
	fmt.Printf("Value of z is %s\n", z)
	fmt.Printf("Value of score is %d\n", score)

	// 2. Print the string value enclosed by double quotes ("Gophers")

	fmt.Printf("Value of z is %q\n", z)

	// 3. Print each variable using the same verb

	fmt.Printf("Value of x is %v\n", x)
	fmt.Printf("Value of y is %v\n", y)
	fmt.Printf("Value of z is %v\n", z)
	fmt.Printf("Value of score is %v\n", score)

	// 4. Print the type of y and score

	fmt.Printf("Type of y is %T\n", y)
	fmt.Printf("Type of y is %T\n", score)
}
