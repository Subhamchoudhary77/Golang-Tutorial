package main

import "fmt"

func main() {

	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a, b, c)

	const (
		a1 = iota
		b2
		c3
	)

	fmt.Println(a1, b2, c3)

	const (
		North = iota //by default 0
		East         //omitting type and value means, repeating its type and value so East = iota = 1 (it increments by 1 automatically)
		South        // -> 2
		West         // -> 3
	)

	// Initializing the constants using a step:
	const (
		c11 = iota * 2 // -> 0
		c22            // -> 2
		c33            // -> 4
	)
	// to skip a particular value
	const (
		c12 = iota * 3 // -> 0
		_              // skipping 3
		c23            // -> 6
		c34            // -> 9
	)

	fmt.Println(c12, c23, c34)

}
