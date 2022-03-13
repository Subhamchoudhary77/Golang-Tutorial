package main

import "fmt"

// 1. Declare a new type type called duration. Have the underlying type be an int.
type duration int

func main() {

	// 2. Declare a variable of the new type called hour using the var keyword
	var hour duration

	// 3. print out the value of the variable hour & print out the type of the variable hour
	fmt.Printf("Value and Type of hour is %d, %T", hour, hour)

	// 4. assign 3600 to the variable hour using the  = operator
	hour = 3600

	// 5. print out the value of hour
	fmt.Printf("\n%d", hour)
}
