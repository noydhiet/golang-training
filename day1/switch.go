package main

import "fmt"

func main() {
	var usia = 34

	switch usia {
	case 10:
		{
			fmt.Println("Remaja")
			fmt.Println("Teenager")
		}
	case 20, 21, 22, 23, 24:
		fmt.Println("Dewasa")
	default:
		fmt.Println("Unable")
	}

	// switch {
	// case (usia > 20) && (usia < 25):
	// 	fmt.Println("Matang")
	// default:
	// 	fmt.Println("Unable")
	// }
}
