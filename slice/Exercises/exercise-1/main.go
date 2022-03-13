package main

import "fmt"

func main() {
	cities := []string{"Delhi", "Mumbai", "Faridabad", "Jaipur", "Noida"}

	for i, v := range cities {
		fmt.Printf("At number %d city is %q \n", i, v)
	}
}
