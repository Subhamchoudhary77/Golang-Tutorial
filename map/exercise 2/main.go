package main

import "fmt"

func main() {
	// 1. Declare a map called m1 which value is nil. Print out its type and value.
	var m1 map[int]int

	fmt.Printf("Type of map: %T and Value of map: %#v\n", m1, m1)

	// 2. Declare a map called m2. It's keys are of type int and values of type string.  Initialize the map using  a map literal with two key:value pairs.

	m2 := map[int]string{}

	// 3. Add the following key: value to the map: 10: "Abba"

	m2[10] = "Abba"

	// 4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.

	fmt.Println(m2[10])
	fmt.Println(m2[1])
}
