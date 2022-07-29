package main

func main() {

	// Problem-1

	//** There are ONLY boolean constants, rune constants, integer constants, **//
	//** floating-point constants, complex constants, and string constants. **//

	// const x int = 10

	// declaring a constant of type slice int ([]int)
	// ERROR -> const initializer []int literal is not a constant
	// const m = []int{1: 3, 4: 5, 6: 8}
	// _ = m

	// Problem-2

	// ERROR -> invalid operation: a * b (mismatched types int and float64)
	// const a int = 7
	// const a = 7 // make `a` untyped constant
	// const b float64 = 5.6
	// const c = a * b

	// x := 8
	// _ = x
	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const xc int = x  // variables belong to runtime

	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const noIPv6 = math.Pow(2, 128) // functions calls belong to runtime
}
