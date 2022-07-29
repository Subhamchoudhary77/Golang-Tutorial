package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: false, 3: false}

	// 1. Delete a key:value pair from the map.

	delete(m, 1)

	fmt.Println(m)

	// 2.  Iterate over the map and print out each key and value.

	for k, v := range m {
		fmt.Printf("Key: %d Value: %v\n", k, v)
	}
}
