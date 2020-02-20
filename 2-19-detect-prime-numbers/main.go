package main

import (
	"strconv"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------

func main() {
	const (
		usage = "Usage: provide as many [agruments] as you want"
		
	)
	var (
		value int
		err error
		x = 5
		y = 2
	)
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println(usage)
		return
	}
	
	next:
	for _, v := range args {
		if value, err = strconv.Atoi(v); err != nil {
			continue
		} 

		switch {
		case value == 2 || value == 3: 
			fmt.Print(value, " ")
			continue
		case value % 2 == 0 || value % 3 == 0:
			continue
		}

		for  x*x <= value {
			if value % x == 0 {
				continue next
			}
			x+= y
			y = 6-y
		}
		fmt.Print(value, " ")
	}

}
