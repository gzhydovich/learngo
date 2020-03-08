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
		for z := -1; z<=7; z++ {
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

			if s2%2 == 0 {
				clock[2][1], clock[2][3], clock[5][1], clock[5][3] = " ░ ", " ░ ", " ░ ", " ░ "
			}

			
			for y := range clock[0] {
				for i := range clock {
					for w := 0; w<=z; w++ {
						if w <= 7 {
							clock[w] = none
						}
					}
					
					fmt.Printf("%v ", clock[i][y])
				}
				fmt.Println()
			}
			time.Sleep(1 * time.Second)
			screen.Clear()	
		}
	}
}