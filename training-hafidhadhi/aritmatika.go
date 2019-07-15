package main

import "fmt"

func main() {

	var value = ((2 + 6) / 2)
	var isEqual = (value == 3)
	fmt.Printf("Nilai %d (%t)", value, isEqual)

	if value == 4 {
		fmt.Println("Match")
	} else if value < 0 {
		fmt.Println("Minus")
	} else {
		fmt.Println("Exception")
	}

}
