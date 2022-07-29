package main

import "fmt"

func main() {
	const (
		//iota starts from zero
		Jun = iota + 6
		Jul //automatically incremented by one
		Aug
	)

	fmt.Println(Jun, Jul, Aug)
}
