package main

import (
	"github.com/inancgumus/screen"
	"fmt"
	"time"
)

var (
	zero = [5]string{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one = [5]string{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two = [5]string{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three = [5]string{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four = [5]string{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five = [5]string{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six = [5]string{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven = [5]string{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight = [5]string{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine = [5]string{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	dots = [5]string{
		" ",
		"█",
		" ",
		"█",
		" ",
	}
	
	nodots = [5]string{
		" ",
		" ",
		" ",
		" ",
		" ",
	}

	digits = [10][5]string{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}


	h1, h2, m1, m2, s1, s2 int
)

func main() {

	for {
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

		clock := [8][5]string{
			digits[h1],
			digits[h2],
			dots,
			digits[m1],
			digits[m2],
			dots,
			digits[s1],
			digits[s2],
		}

		if s2 == 0 || s2%2 == 0 {
			clock[2][1], clock[2][3], clock[5][1], clock[5][3] = "░", "░", "░", "░"
		}

		for y := range clock[0] {
			for i := range clock {
				fmt.Printf("%v ", clock[i][y])
			}
			fmt.Println()
		}

		time.Sleep(1 * time.Second)

		screen.Clear()
	}

	
}
