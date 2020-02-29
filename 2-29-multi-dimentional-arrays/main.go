package main

import (
	"time"
	"math/rand"
	"fmt"
	"os"
)

const (
	usage = "Usage: [your name] [positive/negative]"
)
func main()  {
	args := os.Args[1:]
	moods := [...][3]string{{"sad", "bad", "awful"}, {"happy", "awesome", "outstanding"}}

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	rand.Seed(time.Now().UnixNano())

	switch args[1] {
	case "negative":
		fmt.Printf("%s is %s", args[0], moods[0][rand.Intn(len(moods[0]))])
	case "positive":
		fmt.Printf("%s is %s", args[0], moods[1][rand.Intn(len(moods[0]))])
	default:
		fmt.Println(usage)
		return
	}

} 