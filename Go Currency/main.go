package main

import (
	"fmt"
	"log"

	"github.com/Rhymond/go-money"
)

func main() {

	// Initialization
	// Initialize Money by using smallest unit value (e.g 100 represents 1 pound). Use ISO 4217 Currency Code to set money Currency. Note that constants are also provided for all ISO 4217 currency codes.

	pound := money.New(100, money.GBP)
	twoPounds := money.New(200, money.GBP)
	twoEuros := money.New(200, money.EUR)

	// Comparisons must be made between the same currency units.

	result, err := pound.GreaterThan(twoPounds) // false, nil

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	result2, err := pound.LessThan(twoPounds) // true, nil
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result2)

	result3, err := twoPounds.Equals(twoEuros) // false, error: Currencies don't match
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result3)

	result6, err := twoPounds.Equals(twoPounds) // true, nil
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result6)

	result4, err := twoPounds.GreaterThanOrEqual(pound)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result4)

	result5, err := pound.LessThanOrEqual(twoPounds)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result5)

	// Asserts

	// 1. To assert if Money value is equal to zero use IsZero()
	pound = money.New(100, money.GBP)
	result7 := pound.IsZero() // false
	fmt.Println(result7)

	// 2. To assert if Money value is more than zero use IsPositive()
	result8 := pound.IsPositive() // true
	fmt.Println(result8)

	// 3. To assert if Money value is less than zero use IsNegative()
	result9 := pound.IsNegative() // false
	fmt.Println(result9)

	// Operations: Add, Subtract, Multiply, Absolute & Negaetive
	pound2 := money.New(100, money.GBP)
	twoPounds2 := money.New(200, money.GBP)

	// Add
	result10, err := pound2.Add(twoPounds2) // £3.00, nil
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", result10.Display())

	// Subtract
	result12, err := pound2.Subtract(twoPounds2) // -£1.00, nil
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", result12.Display())

	// multiply
	result11 := pound2.Multiply(2) // £2.00
	fmt.Printf("%v\n", result11.Display())

	pound3 := money.New(-100, money.GBP)

	// Absolute
	result13 := pound3.Absolute() // £1.00
	fmt.Printf("%v\n", result13.Display())

	//Negative
	pound4 := money.New(100, money.GBP)
	result14 := pound4.Negative() // £1.00
	fmt.Printf("%v\n", result14.Display())

	// Split

	// In order to split Money for parties without losing any pennies due to rounding differences, use Split().

	pound = money.New(100, money.GBP)
	parties, err := pound.Split(3)

	if err != nil {
		log.Fatal(err)
	}

	parties[0].Display() // £0.34
	parties[1].Display() // £0.33
	parties[2].Display() // £0.33
	fmt.Printf("%v\n", parties[0].Display())
	fmt.Printf("%v\n", parties[1].Display())
	fmt.Printf("%v\n", parties[2].Display())

	// Allocation

	// To perform allocation operation use Allocate().

	// It splits money using the given ratios without losing pennies and as Split operations distributes leftover pennies amongst the parties with round-robin principle.

	pound = money.New(100, money.GBP)
	// Allocate is variadic function which can receive ratios as
	// slice (int[]{33, 33, 33}...) or separated by a comma integers
	parties1, err := pound.Allocate(33, 33, 33)

	if err != nil {
		log.Fatal(err)
	}

	parties1[0].Display() // £0.34
	parties1[1].Display() // £0.33
	parties1[2].Display() // £0.33

	// To format and return Money as a string use Display().

	result20 := money.New(123456789, money.EUR).Display() // €1,234,567.89
	fmt.Printf("%v", result20)
}
