package main

import "fmt"

func main() {

	var value = ((2 + 6) / 2.0)
	var isEqual = (value == 9)
	fmt.Println(value)

	// Manual method
	if value == 4 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// Go method
	fmt.Println("Nilai", value, isEqual)

	if value == 4 {
		fmt.Println("Match")
	} else if value < 0 {
		fmt.Println("Minus")
	} else {
		fmt.Println("Exception")
	}
}
