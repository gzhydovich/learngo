package main



import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)


// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------


const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	messageWin := [3]string{"Hell yeah!", "EZ", "you are the best!"}
	messageLoose := [3]string{"You suck!", "Looser!", "Play some other game"}
	x := rand.Intn(len(messageLoose))
	var n int

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("One or both arguments are incorrect. Pick 2 numbers")
		return
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive numbers.")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {

		if guess1 > guess2 {
			n = rand.Intn(guess1 + 1)
		} else {
			n = rand.Intn(guess2 + 1)
		}
		

		if n != guess1 && n != guess2 {
			continue
		}

		if turn == 1 {
			fmt.Println(messageWin[x])
		} else {
			fmt.Println(messageWin[x])
		}
		return
	}

	fmt.Println(messageLoose[x])
}