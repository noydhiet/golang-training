package main

import "fmt"

func main()  {
	var value = ((2+6)/2)	
	fmt.Println(value)

	// pengganti if
	var isEqual = (value==4)
	fmt.Println("Nilai", value, isEqual)

	if (value == 4) {
		fmt.Println("Match")
	} else if value < 0 {
		fmt.Println("Minus")
	} else {
		fmt.Println("Exception")
	}



}


