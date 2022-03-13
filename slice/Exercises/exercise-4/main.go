// Problem Statement: Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at command line.

// The user should give between 2 and 10 numbers.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please run the program with arguments (2-10 numbers)!")
	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0.0, 1.0

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please provide the arguments between 2 to 10")
	} else {
		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum = sum + num
			product *= num
		}
	}
	fmt.Printf("Value of sum = %f\n", sum)
	fmt.Printf("Value of product = %f", product)
}
