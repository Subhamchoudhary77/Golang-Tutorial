package main

import "fmt"

func main() {
	const books int = 5 //typed constant

	const year = 2021 //untyped constant

	//declaring constant in a group or mulpiple constant

	const (
		a = 2
		b = 3
		c = 4
	)

	fmt.Println(a, b, c)

	const (
		val  = 500
		val1 // val1 and val2 get type and value from first constant i.e. val
		val2
	)

	fmt.Println(val, val1, val2)

	// CONSTANTS RULES

	// 1. You cannot change a constant
	const temp int = 100
	// temp = 50 //compile-time error

	// 2. You cannot initiate a constant at runtime (constants belong to compile-time)
	// const power = math.Pow(2, 3) //error, functions calls belong to runtime

	// 3. You cannot use a variable to initialize a constant
	t := 5
	// error, variables belong to runtime and you cannot initialize a const to runtime values
	// const tc = t

	// 4. You can use a function like len() to initialize a const if it has as argument
	// a constant string literal (not a variable)
	// a string literal is an untyped constant

	const l1 = len("Hello") //ok

	str := "Hello"
	// const l2 = len(str) //error, str is a variable and belongs to runtime

	_, _ = t, str
}
