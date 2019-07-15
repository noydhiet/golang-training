package main

import "fmt"

func main() {

	var usia = 8
	switch {
	case (usia > 25):
		fmt.Println("matang")
	// case 20, 21, 22, 23, 24:
	// 	fmt.Println("dewasa")
	default:
		fmt.Println("Unable")
	}
}
