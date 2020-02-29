package main

import (
	"fmt"
	"strings"
)

func main() {

	books := [3]string{"Harry potter", "Ray allen", "Poor dad, rich dad"}

	upper := books
	lower := books

	for i, v := range books {
		upper[i] = strings.ToUpper(v)
	}

	for i, v := range books {
		lower[i] = strings.ToLower(v)
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("lower: %q\n", lower)
	fmt.Printf("upper: %q\n", upper)
}