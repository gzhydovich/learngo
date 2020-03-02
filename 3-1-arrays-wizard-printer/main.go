package main

import (
	"strings"
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---

func main()  {
	scientists := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "Time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
	}

	for i := range scientists {
		n:= scientists[i]
		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}
}