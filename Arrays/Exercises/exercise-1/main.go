package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1. Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
	var cities [2]string
	// Print out the cities array and notice the value of its elements
	fmt.Printf("%#v\n", cities)

	fmt.Println(strings.Repeat("#", 20))
	// 2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	grades := [3]float64{5.5, 6.5}
	fmt.Printf("%#v\n", grades)

	fmt.Println(strings.Repeat("#", 20))

	// 3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool.
	taskDone := [...]bool{true, false, false, true}
	fmt.Printf("%#v\n", taskDone)

	// 4. Iterate over the array called cities using the classical for loop syntax and len function and print out element by element.
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// 5. Iterate over grades using the range keyword and print out element by element.
	for i, v := range grades {
		fmt.Println(i, v)
	}

}
