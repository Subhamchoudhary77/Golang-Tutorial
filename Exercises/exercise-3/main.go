package main

import "fmt"

// ERROR -> := is not allowed in package scope (only local scope)
// version := "3.1"

var version = "3.1"

func main() {
	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false

	_, _, _ = a, x, y

	fmt.Println(a, x, y)

	name := "Golang"
	fmt.Println(name)
}
