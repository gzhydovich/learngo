package main

import (
	"github.com/inancgumus/screen"
	"fmt"
	"time"
)

func main() {

	var h1, h2, m1, m2, s1, s2 int

	screen.Clear()

	for shift := 0; ; shift++ {
		screen.Clear()	
		screen.MoveTopLeft()
		

		h, m, s := time.Now().Clock()

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
		}

		for line := range clock[0] {
			l := len(clock)

			s, e := shift%l, l

			if shift%(l*2) >= l {
				s, e = 0, s
			}

			for j := 0; j < l-e; j++ {
				fmt.Print("     ")
			}

			for i := s; i < e; i++ {
				next := clock[i][line]
				fmt.Print(next, "  ")

			}
			fmt.Println()
			
		}
	time.Sleep(1 * time.Second)
	}
}