package main

import "fmt"

func main() {
	// 1.  Declare a, b, c, d on a single line for better readability -> multiple declarations

	var (
		a int
		b float64
		c bool
		d string
	)

	// 2. Declare x, y and z on a single line -> multiple short declarations

	x, y, z := 20, 15.5, "Subham"

	// 3. Change the program to run without error using the blank identifier (_)

	_, _, _, _ = a, b, c, d
	_, _, _ = x, y, z

	fmt.Println(a, b, c, d)
	fmt.Println(x, y, z)
}
