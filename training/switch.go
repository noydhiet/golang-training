package main

import "fmt"

func main(){

	var usia = 25

	switch {
	case (usia>=20) && (usia<=25):
		fmt.Println("Dewasa")
	
	default:
		fmt.Println("Tidak diketahui")
	}	
}