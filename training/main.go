package main

import "fmt"

/*

Bla bla bla

*/

func main(){

	// lock the data type of this variables into string
	firstName := "Andhika"

	var lastName string = "Permana"
	address := new(string)
	// this line of code produce Hello World as result
	// without using ane parameter as input
	// fmt.Println("Hello ", firstName, lastName)
	fmt.Printf("Hello %s %s !!!", firstName, lastName)
	fmt.Printf(*address)
} //end of code