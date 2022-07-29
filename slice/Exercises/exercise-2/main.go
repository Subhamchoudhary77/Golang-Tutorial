package main

import "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	// ERROR -> index out of range [2] with length 2
	mySlice[1] = 6

	a := 10
	// mismatch type
	mySlice[0] = float64(a)

	// ERROR -> index out of range [3] with length 2
	mySlice[1] = 10.10

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

}
