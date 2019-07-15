package main

import "fmt"

func main() {
	var usia = 22

	switch {
	case (usia == 10):
		{
			fmt.Println("Remaja")
			fmt.Println("Tanggung")
		}
	case (usia >= 20) && (usia <= 24):
		fmt.Println("Dewasa")
	case (usia > 25):
		fmt.Println("Matang")
	default:
		fmt.Println("Unable")
	}

}
