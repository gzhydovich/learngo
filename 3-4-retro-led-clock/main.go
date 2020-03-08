package main

import (
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

	clock = [8][5]string{}

	h1, h2, m1, m2, s1, s2 int
)

func main() {
	showDots := true

	for y := range digits[0] {
		for i := range digits {
			fmt.Printf("%v ", digits[i][y])
		}
		fmt.Println()
	}
	
	fmt.Println()

	// for {
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

		clock = [8][5]string{digits[h1],digits[h2],digits[m1],digits[m2],digits[s1],digits[s2]}

		if showDots {
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][0], clock[1][0], dots[0], clock[2][0], clock[3][0], dots[0], clock[4][0], clock[5][0])
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][1], clock[1][1], dots[1], clock[2][1], clock[3][1], dots[1], clock[4][1], clock[5][1])
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][2], clock[1][2], dots[2], clock[2][2], clock[3][2], dots[2], clock[4][2], clock[5][2])
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][3], clock[1][3], dots[3], clock[2][3], clock[3][3], dots[3], clock[4][3], clock[5][3])
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][4], clock[1][4], dots[4], clock[2][4], clock[3][4], dots[4], clock[4][4], clock[5][4])
			showDots = false
		} else {
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][0], clock[1][0], nodots[0], clock[2][0], clock[3][0], nodots[0], clock[4][0], clock[5][0])
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][1], clock[1][1], nodots[1], clock[2][1], clock[3][1], nodots[1], clock[4][1], clock[5][1])
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][2], clock[1][2], nodots[2], clock[2][2], clock[3][2], nodots[2], clock[4][2], clock[5][2])
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][3], clock[1][3], nodots[3], clock[2][3], clock[3][3], nodots[3], clock[4][3], clock[5][3])
			fmt.Printf("%s %s %s %s %s %s %s %s\n", clock[0][4], clock[1][4], nodots[4], clock[2][4], clock[3][4], nodots[4], clock[4][4], clock[5][4])
			showDots = true
		}

		
		time.Sleep(1 * time.Second)
	// }
	

}
