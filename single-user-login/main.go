package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	const username, password = "jack", "6475"
	var (
		args = os.Args
		l = len(args)
	)

	if l != 3 {
		fmt.Println(" Usage: [username] [password]")
		return
	} 

	if args[1] == username {
		if args[2] == password {
			fmt.Printf("Access granted to %q.", args[1])
		} else {
			fmt.Printf("Invalid password for %q.", args[1])
		}
	} else {
		fmt.Printf("Access denied for %q.", args[1])
	}
}