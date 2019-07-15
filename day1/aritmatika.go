package main

import "fmt"

func main() {
	var value = ((2 + 6) / 2)
	var isEqual = (value == 4)
	fmt.Printf("Nilai %d (%t)", value, isEqual)
	fmt.Println()

	if value == 4 {
		fmt.Println("Match")
	} else if value < 0 {
		fmt.Println("Minus")
	} else {
		fmt.Println("Exception")
	}

}
