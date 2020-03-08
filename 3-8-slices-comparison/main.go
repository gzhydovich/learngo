package main

import (
	"sort"
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesASlice := strings.Split(namesA, ", ")

	sort.Strings(namesASlice)
	sort.Strings(namesB)

	for i, v := range namesASlice {
		if v != namesB[i] {
			fmt.Println("Not Eq")
			return
		}
	}

	if len(namesB) == len(namesASlice) {
		fmt.Println("Eq")
	}
}