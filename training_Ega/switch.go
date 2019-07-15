package main

import "fmt"

func main() {

	var usia int = 27

	switch {
	// case 10:
	// 	{
	// 		fmt.Println("Remaja")
	// 		fmt.Println("Tanggung")
	// 	}
	// case 20, 21, 22, 23, 24:
	// 	fmt.Println("Dewasa")
	case usia > 25 && usia < 30:
		fmt.Println("Matang")
	default:
		fmt.Println("Unable")
	}
}
