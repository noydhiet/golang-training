package main

import "fmt"

func main()  {
	var usia = 22

	switch {
	case (usia > 20) && (usia < 25):
		fmt.Println("Matang")
	default:
		fmt.Println("Unable")
	}
}