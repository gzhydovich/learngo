package main

import (
	"fmt"
	"bytes"
)

// ---------------------------------------------------------
// EXERCISE: Append
//
//  Please follow the instructions within the code below.
//
// EXPECTED OUTPUT
//  They are equal.
//
// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal
// ---------------------------------------------------------

func main() {
	// 1. uncomment the code below
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, png...)

	ok := bytes.Equal(png, header)

	if !ok {
		fmt.Println("They are NOT equal")
		return
	}

	fmt.Println("They are equal")
}