package main

// using operator aritmatika to do some calculation

import "fmt"

func main() {
	var value = ((2 + 6) / 2)
	// var isEqual = (value == 4)
	// fmt.Println("Nilai", value, isEqual)

	if value == 4 {
		fmt.Println("match")
	} else if value < 0 {
		fmt.Println("minus")
	} else {
		fmt.Println("exception")
	}
}
