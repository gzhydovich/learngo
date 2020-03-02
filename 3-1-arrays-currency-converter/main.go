package main

import (
	"strconv"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------

func main() {
	const (
		usage = "Usage:  Please provide the amount to be converted."

		EUR = iota
		GBP
		JPY
	)
	
	rates := [...]float64{
		EUR: 2,
		GBP: 2.5,
		JPY: 1.2,
	}

	arg := os.Args[1:]

	var (
		amount float64
		err error
	)

	if len(arg) != 1 {
		fmt.Println(usage)
		return
	}

	if amount, err = strconv.ParseFloat(arg[0], 64); err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", amount, rates[EUR] * amount)
	fmt.Printf("%.2f USD is %.2f GBP\n", amount, rates[GBP] * amount)
	fmt.Printf("%.2f USD is %.2f JPY\n", amount, rates[JPY] * amount)
}