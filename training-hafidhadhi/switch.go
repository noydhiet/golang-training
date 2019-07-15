package main

import "fmt"

func main() {

	var usia = 24
	switch usia {
	case 10:
		{
			fmt.Println("Remaja")
			fmt.Println("Tanggung")
		}
	case 20, 21, 22, 23, 24:
		fmt.Println("Dewasa")
	// case usia > 25:
	// 	fmt.Println("Matang")
	default:
		fmt.Println("Unable")
	}

}
