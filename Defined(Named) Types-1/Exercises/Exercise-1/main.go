package main

import "fmt"

type duration int

func main() {
	var hour duration = 3600
	minute := 60
	// mismatch type that's why we explicitly converting
	fmt.Println(int(hour) != minute)
}
