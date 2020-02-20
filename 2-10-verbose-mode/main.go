package main

import (
	"strconv"
	"fmt"
	"os"
	"time"
	"math/rand"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	usage = `Usage: [guess]
-v   to get verbose output`
	usageVerbose = "You've enabled verbose output, but need to provide the guess as well"
	maxTurns = 5
	messageLoose = "You lost"
	messageWon = "You Won!"

)

var (
	guess int
	err error
	verbose bool
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 || len(args) > 2 {
		fmt.Println(usage)
		return
	}

	if args[0] == "-v" {
		verbose = true
	}

	guess, err = strconv.Atoi(args[len(args)-1])

	if err != nil {
		fmt.Print("Error! Arguments are incorrect \n", usage)
		return
	}

	if (len(args) == 2 && args[0] != "-v") || len(args) > 2 {
		fmt.Print(usage)
		return
	} 

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(maxTurns+1)
		if verbose {
			fmt.Print(n, " ")
		}
		if n == guess {
			fmt.Println(messageWon)
			return
		} 
	}
	fmt.Println(messageLoose)


	// if len(args) < 1 || len(args) > 2 {
	// 	fmt.Println(usage)
	// 	return
	// } else if len(args) == 1 {
	// 	if guess, err = strconv.Atoi(args[0]); err != nil {
	// 		fmt.Println(usage)
	// 		return
	// 	} else {
	// 		//program with 1 param
	// 		for turn := 1; turn <= maxTurns; turn++ {
	// 			n := rand.Intn(maxTurns+1)
	// 			if n == guess {
	// 				fmt.Println(messageWon)
	// 				return
	// 			}
	// 		}
	// 	}

	// } else if len(args) == 2 {
	// 	if guess, err = strconv.Atoi(args[1]); err == nil && args[0] == "-v" {
	// 		//Program with 1 param and -v
	// 		for turn := 1; turn <= maxTurns; turn++ {
	// 			n := rand.Intn(maxTurns+1)
	// 			if n == guess {
	// 				fmt.Println(messageWon)
	// 				return
	// 			} 
	// 			fmt.Print(n, " ")
	// 		}
	// 	} else {
	// 		fmt.Println(usage)
	// 		return
	// 	}
	// }
	// fmt.Println(messageLoose)
}