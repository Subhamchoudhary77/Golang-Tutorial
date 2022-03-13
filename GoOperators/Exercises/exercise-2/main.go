package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// 1. i to string (int to string)

	fmt.Printf("Type & Value of of i is %T, %q\n", strconv.Itoa(i), strconv.Itoa(i))

	// 2. s2 to int (string to int)

	new, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Type & Value of new is %T, %d\n", new, new)

	// 3. float64 to string

	ss1 := fmt.Sprintf("%f", f)

	fmt.Printf("Type and Value of ss1 is %T, %s\n", ss1, ss1)

	// 4. 4. s1 to float64  (string to float64)

	conv, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Type & Value of conv is %T, %f", conv, conv)
}
