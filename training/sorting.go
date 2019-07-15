package main

import (
	"fmt"
)

var urutanAcak [6]int = [6]int{10, 3, 2, 1, 8, 7}

func bubbleSort(input [6]int) {
	// n is the number of items in our list
	n := 6
	// set swapped to true
	swapped := true
	// loop
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 1; i < n; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i-1] > input[i] {
				// log that we are swapping values for posterity

				// swap values using Go's tuple assignment
				input[i], input[i-1] = input[i-1], input[i]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}
	// finally, print out the sorted list
	fmt.Println(input)
}

func main() {
	bubbleSort(urutanAcak)
}
