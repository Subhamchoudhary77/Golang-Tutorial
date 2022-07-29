package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	// Integer Overflow

	var x uint8 = 255 // This is maximum value of uint8
	x++               // This is overflow, and go reset its minimum value of uint8 i.e. 0, but go doesn't catches compile time error because go catches for expression which is evaluated at compile time, but here is no expression available

	fmt.Println(x)

	// a := int8(255 + 1) // this is error: constant 256 overflows int8

	var y int8 = 127 // This is maximum value of int8
	y += 1           // This is overflow, and go reset its minimum value of int8 i.e. -128,

	fmt.Printf("%d", y) // -128

	// Float Overflow

	f := float32(math.MaxFloat32) // This is maximum value of float32
	fmt.Println(f)

	f = f * 2
	fmt.Println(f) // f overflows to infinite

	// Note:

	// if you want to use big number, but you know it would overflow, then you should use big package

	// example

	var b int = 500000

	newBigInt := big.NewInt(int64(b))

	fmt.Println(newBigInt)

}
