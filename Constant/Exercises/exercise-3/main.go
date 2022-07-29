package main

import "fmt"

func main() {
	//  Calculate how many seconds are in a year.

	// 	STEPS:

	//  1. Declare secPerDay constant and initialize it to the number of seconds in a day

	const secPerDay int64 = 24 * 60 * 60

	//  2. Declare daysYear constant and initialize it to 365

	const daysYear int64 = 365

	//  3. Use fmt.Printf() to print out the total number of seconds in a year.

	const secPerYear int64 = secPerDay * daysYear // 31536000

	fmt.Printf("Number of seconds in a year = %d", secPerYear)
}
