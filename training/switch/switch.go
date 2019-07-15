package main

import "fmt"

func main() {
	usia := 24

	switch {
	case (usia > 10):
		fmt.Println("Remaja")
	// case 20, 21, 22, 24:
	// 	fmt.Println("Dewasa")
	default:
		fmt.Println("Bayi")
	}
}
