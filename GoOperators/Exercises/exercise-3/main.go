package main

import "fmt"

func main() {
	x, y := 4, 5.1

	// Mismatch Type
	z := float64(x) * y
	fmt.Println(z)

	// Mismatch Type
	a := 5
	b := 6.2 * float64(a)
	fmt.Println(b)

}
