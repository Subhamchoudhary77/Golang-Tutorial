package main

import (
	"fmt"
	"strconv"
)

func main() {

	// s := string(99)

	// fmt.Println(s)

	// Sprintf (in which format you want to get data, like int or float, etc, number which you want to convert)
	// example: Sprintf("%d", 77)

	var myStr = fmt.Sprintf("%f", 44.2)

	fmt.Println(myStr)
	fmt.Printf("%T\n", myStr)

	// To convert string to float

	var s1 = "123.56"

	var s2 = "56"
	// To convert string to float, int, uint and bool, use strconv.ParseFloat, etc.
	f1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f1)

	f2, err := strconv.ParseInt(s2, 0, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f2)
	fmt.Printf("%T\n", f2)

	// Another way to convert string to int

	i, err := strconv.Atoi(s2) // Limitation: you can convert into int only, not in int32, and int64
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	// Let's convert int to string

	var a int64 = 77

	j := strconv.Itoa(int(a))

	fmt.Println(j)
	fmt.Printf("%T\n", j)
}
