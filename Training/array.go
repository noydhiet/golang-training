package main

import "fmt"

func main(){

	var peserta [2] string
	peserta[0] = "Nanang"
	peserta[1] = "Reysy"

	usia := [3] int {1,2,3}

	fmt.Println(peserta[1])
	fmt.Println(usia)
	
	for i :=0; i<3; i++{
	fmt.Println(peserta)
	}
}