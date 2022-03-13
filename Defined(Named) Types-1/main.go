package main

import "fmt"

func main() {
	type age int           // int is its underlying type
	type oldAge age        // int is its underlying type
	type veryOldAge oldAge // int is its underlying type

	type speed int64

	var s1 speed = 20
	var s2 speed = 10

	fmt.Printf("Result after subtracting s1 to s2 = %d", s1-s2)

	fmt.Printf("Type of s1 is %T", s1)
}
