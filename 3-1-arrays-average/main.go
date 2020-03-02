package main

import (
	"strconv"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func main() {
	const (
		usage = "Please tell me numbers (maximum 5 numbers)"
	)
	args := os.Args[1:]
	argsLen := len(args) 
	var sum, count float32
	

	if argsLen > 5 || argsLen == 0 {
		fmt.Println(usage)
		return
	}

	var array [5]float32

	fmt.Print(len(args))

	for i, v := range args {
		if value, err := strconv.ParseFloat(v, 32); err != nil {
			continue
		} else {
			array[i] = float32(value)
			sum += float32(value)
			count++
		}
	}

	fmt.Printf("Your numbers: %v\n", array)
	fmt.Printf("Average: %.1f", sum/count)
}