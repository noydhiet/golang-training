package main

import "fmt"

/*
Main class for hello world
*/

func main() {

	// lock data-type of variable into string
	firstName := "Hafidh"
	var lastName string = "Adhi"
	// this line of code produce hello world as a result
	// without using any parameter as input
	fmt.Println("Hello ", firstName, lastName, "!!!")
	fmt.Printf("Hello %s %s %s!\n", firstName, lastName, "!!!")
	//fmt.Printf(*address)
} // end of code
