package main

import (
	"strings"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

func main() {
	const (
		usage = "Usage: provide params to search"
		s = "we can do it over and over and again and once again"
	)
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println(usage)
		return
	}

	words := strings.Split(s, " ")

	for _, v := range args {
		for i, w := range words {
			if w == strings.ToLower(v) {
				fmt.Printf("%-2d : %s\n", i+1, w)
			}
		}
	}
}