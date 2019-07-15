package main

import "fmt"

func main() {
	var value = 4 * 4
	var isEqual = (value == 16)
	fmt.Println("nilai", value, isEqual)

	if value == 4 {
		fmt.Println("Match")
	} else if value < 0 {
		fmt.Println("Minus")
	} else {
		fmt.Println("ga tw")
	}
}
