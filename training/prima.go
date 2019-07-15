package main

import "fmt"

func main(){

	var input int 
	fmt.Println("Masukan nilai:")
	/*
	buat fungsi masukan 
	*/
	fmt.Scanf("%d", &input)

	temp:= 0

	for i := 1; i <= input; i++ {
		if input%i == 0 {
			temp++
		}
	}
	if temp == 2 {
		fmt.Println("Adalah bilangan prima")
	}else{
		fmt.Println("Bukan bilangan prima")
	}
}


