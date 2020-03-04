package main

import (
	"strings"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Word Finder
//
//   Your goal is to search for the words inside the corpus.
//
//   Note: This exercise is similar to the previous word finder program:
//   https://github.com/inancgumus/learngo/tree/master/13-loops/10-word-finder-labeled-switch
//
//   1. Get the search query from the command-line (it can be multiple words)
//
//   2. Filter these words, do not search for them:
//      and, or, was, the, since, very
//
//      To do this, use an array for the filtered words.
//
//   3. Print the words found.
//
// RESTRICTION
//   + The search and the filtering should be case insensitive
//
// HINT
//   + strings.Fields function converts a given string to a slice.
//
//     You can find its example in the word finder program that I've mentioned
//     above.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me a word to search.
//
//   go run main.go and was
//
//   go run main.go AND WAS
//
//   go run main.go cat beginning
//     #2 : "cat"
//     #11: "beginning"
//
//   go run main.go Cat Beginning
//     #2 : "cat"
//     #11: "beginning"
// ---------------------------------------------------------

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"
const filter = "and or was the since very"
const usage = "Usage: Provide word/s for search"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(usage)
		return
	}

	corpusArr := strings.Fields(corpus)
	filterArr := strings.Fields(filter)


	for i , v := range corpusArr {
		text:
		for _, w := range args {
			for _, y := range filterArr {
				if strings.ToLower(y) == strings.ToLower(w) {
					continue text
				}
			}

			if strings.ToLower(v) == strings.ToLower(w) {
				fmt.Printf("%-2d: %q\n", i+1, v)
			}
		}
	}
}