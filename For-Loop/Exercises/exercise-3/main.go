// Problem Statement: Use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.

package main

import "fmt"

func main() {
	count := 0
	for i := 1; i < 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
			count++
		}
		if count == 3 {
			break
		}
	}
}
