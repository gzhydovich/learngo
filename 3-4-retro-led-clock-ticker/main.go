package main

import (
	"github.com/inancgumus/screen"
	"fmt"
	"time"
)

func main() {

	var h1, h2, m1, m2, s1, s2 int

	screen.Clear()

	for {
		screen.MoveTopLeft()
		h, m, s := time.Now().Clock()
		ns := time.Now().Nanosecond()/100000000

		if h >= 10 {
			h1 = h / 10
			h2 = h % 10
		} else {
			h1 = 0
			h2 = h % 10
		}

		if m >= 10 {
			m1 = m / 10
			m2 = m % 10
		} else {
			m1 = 0
			m2 = m % 10
		}

		if s >= 10 {
			s1 = s / 10
			s2 = s % 10
		} else {
			s1 = 0
			s2 = s % 10
		}

		clock := [...]placeholder{
			digits[h1], digits[h2],
			dots,
			digits[m1], digits[m2],
			dots,
			digits[s1], digits[s2],
			dot,
			digits[ns],
		}

		if s2 == 0 {
			clock = alarm
		}

		if clock != alarm && s2%2 == 0 {
			clock[8][4], clock[2][1], clock[2][3], clock[5][1], clock[5][3] = "░", "░", "░", "░", "░"
		}

		for y := range clock[0] {
			for i := range clock {
				fmt.Printf("%v ", clock[i][y])
			}
			fmt.Println()
		}

		const splitSecond = time.Second / 10
		time.Sleep(splitSecond)
		screen.Clear()
	}
}