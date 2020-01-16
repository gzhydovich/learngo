package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
	var (
		args = os.Args
		l = len(args) - 1
	)

	if l != 1 || len(args[1]) != 1 {
		fmt.Print("Give me letter argument")
		return
	}

	var (
		number = strings.IndexAny(os.Args[1], "1234567890") == 0
		vowel = strings.IndexAny(os.Args[1], "aeiou") == 0
		canBeBoth =  strings.IndexAny(os.Args[1], "yw") == 0
	)

	if vowel {
		fmt.Printf("%q is a vowel.", os.Args[1])
	} else if canBeBoth {
		fmt.Printf("%q is sometimes a vowel, sometimes not.", os.Args[1])
	} else if number {
		fmt.Printf("%q is not a letter.", os.Args[1])
	} else {
		fmt.Printf("%q is a consonant.", os.Args[1])
	}
}