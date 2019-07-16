package main

import (
	"fmt"
	"os"
	"strings" 
)

func main()  {
	
	var usia=20

	switch {

	case (usia > 25): {
		fmt.Println("Remaja")
	default:
		fmt.Println("Unable")
	}


}