package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------


func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs                       []string
		sizes, beds, baths, prices []int
		averages []float64
		sizesSum, bedsSum, bathsSum, pricesSum int
	)

	rows := strings.Split(data, "\n")
	

	for _, row := range rows {
		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)
		sizesSum+=n

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)
		bedsSum+=n

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)
		bathsSum+=n

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
		pricesSum+=n
	}

	len := len(rows)
	averageSum := float64(sizesSum)/float64(len)
	averageBeds := float64(bedsSum)/float64(len)
	averageBaths := float64(bathsSum)/float64(len)
	averagePrices := float64(pricesSum)/float64(len)
	averages = append(averages, averageSum, averageBeds, averageBaths, averagePrices)

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	fmt.Printf("%-15s", "")
	fmt.Printf("%-15.2f", averageSum)
	fmt.Printf("%-15.2f", averageBeds)
	fmt.Printf("%-15.2f", averageBaths)
	fmt.Printf("%-15.2f", averagePrices)
}