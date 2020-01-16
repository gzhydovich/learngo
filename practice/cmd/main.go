package main

import (
	"strings"
	"unicode/utf8"
	"strconv"
	"path"
	"time"
	"fmt"
	"os"
)
	
func main() {

	var speed int

	speed = 100
	fmt.Println(speed)

	/////////////////////////////////

	speed = speed - 25
	fmt.Println(speed)

	//////////////////////////////////


	var (
		speed1 int
		now time.Time //declared a variable by pointing to a package type
	)

	speed1, now = 100, time.Now()
	fmt.Println(speed1, now)

	////////////////////////////////////

	var (
		speed2 = 100
		prevSpeed = 50
	)

	speed2, prevSpeed = prevSpeed, speed2

	///////////////////////////////////// methods with multiple outputs

	_, file := path.Split("gene/file.txt")

	fmt.Println("file:", file)

	//////////////////////////////////// Convertion of values

	speed3 := 100
	force := 2.5

	speed3 = int(float64(speed3) * force)

	fmt.Println(speed3)

	///////////////////////////////////// convert byte to string results in letter codes convertion

	fmt.Println(string([]byte{104, 105}))

	///////////////////////////////////// Printf gives you type as well

	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])


	fmt.Println("Number of items in os.Args",len(os.Args))

	/////////////////////////printf

	var brand = "Google"

	fmt.Printf("%q\n", brand)
	fmt.Printf("%s\n", brand)

	var ops = 2350
	var ok = 543
	var fail = 433

	fmt.Printf("total: %d success: %d / %d\nImagine %[2]d/%[2]d \n", ops, ok, fail)


	fmt.Println("hi \n\"hi\"")

	var heat float64

	fmt.Printf("%T\n", heat)

	///////////////////// convert integer and bool to a string

	eq := "1 + 2 = "
	sum := 1+2
	
	fmt.Println(eq + strconv.Itoa(sum))
	fmt.Println(strconv.FormatBool(false) + " string")

	//////////////////// len - returns bytes

	fmt.Println(len(eq))

	fmt.Println(utf8.RuneCountInString(eq))

	//////////////// String exercise

	msg := os.Args[1]

	l := len(msg)

	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)

	fmt.Println(s)

	/////////// Types

	type Duration int64

	var ms int64 = 1000
	var ns Duration

	ns = Duration(ms)

	fmt.Println(ns)
	 

}