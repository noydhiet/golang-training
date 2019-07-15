package main

import "fmt"


func main(){

	var a int = 10
	multipleA(&a, 10)
	printA(a)
}

func multipleA (a *int, b int){
	*a = *a * b
}

func printA( a int){
	fmt.Println(a)
}