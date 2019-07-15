package main

import "fmt"

func main() {
	var value = ((2.0 + 6) / 2)
	var isEqual = (value == 4)
	fmt.Println("Nilai", value, isEqual)

	if value == 4 {
		fmt.Println("Match")
	} else if value < 0 {
		fmt.Println("Minus")
	} else {
		fmt.Println("Exception")
	}
}