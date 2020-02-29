package main

import (
	"time"
	"math/rand"
	"fmt"
	"os"
)

const (
	usage = "Usage: [your name]"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	moods := [4]string{"angry", "happy", "excited", "odd"}
	
	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	fmt.Printf("%s feels %s\n", args[0], moods[rand.Intn(len(moods))])
	fmt.Printf("%#v\n", moods)
	fmt.Printf("%q\n", moods)
	fmt.Printf("%T", moods)

}