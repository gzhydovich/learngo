package main

import "fmt"

func main() {
	var (
		names     [3]string  // The names of your best three friends
		distances [5]int     // The distances to five different locations
		data      [5]byte    // A data buffer with five bytes of capacity
		ratios    [1]float64 // Currency exchange ratios only for a single currency
		// zero      [0]byte    // A byte array that doesn't occupy memory space
	)


	names[0] = "Dan"
	names[1] = "Ken"
	names[2] = "Ban"

	distances = [5]int{1, 2, 4, 13, 7}

	data = [5]byte{'F', 'u', 'c', 'k', '!'}

	ratios = [1]float64{3.14145}

	alives := [...]bool{true, false, true, true}


	for i := 0; i<=len(names)-1; i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	for i, v :=  range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	for i, v :=  range data {
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	for i, v :=  range ratios {
		fmt.Printf("ratios[%d]: %v\n", i, v)
	}

	for i, v :=  range alives {
		fmt.Printf("alives[%d]: %v\n", i, v)
	}

///////////////////////////////////////////////
	namess := [3]string{"Einstein", "Shepard", "Tesla"}

	books := [5]string{"Kafka's Revenge", "Stay Golden"}

	fmt.Printf("%q\n", namess)
	fmt.Printf("%q\n", books)

/////////////////////////////////////////////

	week := [...]string{"Monday", "Tuesday"}
	wend := [...]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int32{2, 4, 6, 8, 10}
	evens2 := [...]int32{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]byte{'h', 'i'}
	var datas [5]uint8 

	fmt.Println(datas == image)
}