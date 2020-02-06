package main

import (
	"strconv"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
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
		max int
		min int
		err error
		sum int
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
		sum+=i
		fmt.Print(i)
		if i != max {
			fmt.Print(" + ")
			continue
		}
		fmt.Printf(" = %d", sum)
	}
}