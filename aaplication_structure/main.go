package main

import "fmt"

func main() {
	SayHello()
}

// SayHello method
func SayHello() {
	fmt.Println("Hello package v1.0.0!")

	var x int = 10
	_ = x
}
