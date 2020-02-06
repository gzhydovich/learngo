package main

import (
	"os"
	"strconv"
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	const (
		errArgs = "Usage: [min] [max]"
	 	errMin = "Incorrect min value provided: %q"
		errMax = "Incorrect max value provided: %q"
		errMaxMin = "Error: Max can't be lower than Min"
	)

	var (
		args = os.Args[1:]
		max, min, sum int
		err error
	)

	if len(args) != 2 {
		fmt.Println(errArgs)
		return
	}

	if min, err = strconv.Atoi(args[0]); err != nil {
		fmt.Printf(errMin, args[0])
		return
	} else if max, err = strconv.Atoi(args[1]); err != nil {
		fmt.Printf(errMax, args[1])
		return
	} else if max < min {
		fmt.Print(errMaxMin)
		return
	}

	for i := min; i <= max; i++ {
		if i % 2 != 0 {
			continue
		}
		sum+=i
		fmt.Print(i)
		if i != max {
			fmt.Print(" + ")
			continue
		}
		fmt.Printf(" = %d", sum)
	}
}