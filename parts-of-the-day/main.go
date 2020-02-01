package main

import (
	"time"
	"fmt"
)

func main() {
	switch h := time.Now().Hour(); {
	case  h >= 6:
		fmt.Println("Good Morning")
	case h >= 12: 
		fmt.Println("Good Afternoon")
	case h >= 18:
		fmt.Println("Good Evening")
	case h < 6:
		fmt.Println("Good Night")
	}
}