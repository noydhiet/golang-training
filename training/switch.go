package main

import "fmt"

func main() {
	var usia = 21

	switch usia {
	// case 10:
	// 	{
	// 		fmt.Println("Remaja")
	// 		fmt.Println("tanggung amat bre")
	// 	}
	// case 20:
	// 	fmt.Println("Dewasa")
	case (usia > 25) && (usia < 30):
		fmt.Println("setengah matang")
	default:
		fmt.Println("Unable")
	}
}
