package main

import "fmt"

func main() {
	var usia = 28

	switch {
	case (usia > 25):
		fmt.Println("Matang")
	default:
		fmt.Println("Unable")
	}

}
