package main

import "fmt"

func main(){

	var usia = 30

	switch{
	case (usia > 20) && (usia < 25):
		fmt.Println("Matang")
	case (usia > 25):
		fmt.Println("Dewasa banget")
	default:
		fmt.Println("Unable")
	}

	}
