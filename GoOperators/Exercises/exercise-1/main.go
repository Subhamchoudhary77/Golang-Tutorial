package main

import "fmt"

func main() {
	var i int = 3
	var f float64 = 3.2

	// converting i to float64

	fmt.Printf("Type & Value of i after converting is %T, %f\n", float64(i), float64(i))

	// converting f to int

	fmt.Printf("Type & Value of f after converting is %T, %d\n", int(f), int(f))

}
