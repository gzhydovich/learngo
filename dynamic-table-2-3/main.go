package main

import (
	"strconv"
	"fmt"
	"os"
)

func main() {
	const (
		usage = "Usage: [table size]"
	)
	var args = os.Args

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	if size, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Provided argument can't be converted to int")
	} else {
		fmt.Printf("%5s", "X")
		for i:= 1; i <= size; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()
		for i := 1; i <= size; i++ {
			fmt.Printf("%5d", i)
			for j := 1; j <= size; j++ {
				fmt.Printf("%5d", j*i)
			}
			fmt.Println()
		}
	}

}