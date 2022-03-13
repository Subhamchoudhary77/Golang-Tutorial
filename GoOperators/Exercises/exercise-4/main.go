package main

import "fmt"

func main() {

	const (
		distanceFromSun = 149_600_000 * 1000 // distance from the Sun to Earth in m
		// (it's allowed to use underscores in numbers for a better readability)

		speedOfLight = 299_792_458 // speed of light in m / s
	)

	const time = distanceFromSun / speedOfLight

	fmt.Printf("Time takes by Sunlight to reach the Earth is %v seconds.", time)

}
