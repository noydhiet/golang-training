package main

import "fmt"

func main() {

	var usia = 26

	switch {
	case usia == 10:
		{
			fmt.Println("Remaja")
			fmt.Println("Tanggung")
		}
	case (usia > 25) && (usia < 30):
		fmt.Println("Matang")
	default:
		fmt.Println("Unable")
	}

}
