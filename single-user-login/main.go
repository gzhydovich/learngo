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
	const (
		usage = "Usage: [username] [password]"
		errUser = "Access denied for %q."
		errPass = "Invalid password for %q."
		accessOK = "Access granted to %q."
	)
	var (
		// args = os.Args
		// l = len(args)

		usernames = [2]string{"jack", "gene"}
		passwords = [2]string{"6475", "1234"}
	)

	// if l != 3 {
	// 	fmt.Println(usage)
	// 	return
	// }

	// if args[1] == usernames[0] || args[1] == usernames[1] {
	// 	if (args[1] == usernames[0] && args[2] == passwords[0]) || (args[1] == usernames[1] && args[2] == passwords[1]) {
	// 		fmt.Printf(accessOK, args[1])
	// 	} else {
	// 		fmt.Printf(errPass, args[1])
	// 	}
	// } else {
	// 	fmt.Printf(errUser, args[1])
	// }

	if args := os.Args; len(args) != 3 {
		fmt.Println(usage)
	} else if args[1] == usernames[0] || args[1] == usernames[1] {
		if (args[1] == usernames[0] && args[2] == passwords[0]) || (args[1] == usernames[1] && args[2] == passwords[1]) {
			fmt.Printf(accessOK, args[1])
		} else {
			fmt.Printf(errPass, args[1])
		}
	} else {
		fmt.Printf(errUser, args[1])
	}
}