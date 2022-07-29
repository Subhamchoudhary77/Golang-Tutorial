// Problem Statement: Use the continue statement to print out all the numbers divisible by 7 between 1 and 50.

package main

import "fmt"

func main() {
	for i := 1; i < 50; i++ {
		if i%7 != 0 { // number which is not divisible by 7
			continue // skip all numbers which is not divisible by 7
		}
		fmt.Println(i)
	}
}
