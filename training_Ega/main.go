package main

import "fmt"

/*
Main class for hello world
*/

func main() {

	firstName := "Muhammad" // lock data type of this var into string
	var lastName string = "Ega"
	var num int = 3
	// this line of code produce Hello World as result
	// without using any parameter as input
	fmt.Println("Hello", firstName, lastName, "!!!")
	fmt.Printf("Hello %s %s !!!\n", firstName, lastName)
	fmt.Printf("Alamatnya : ", num)
} // end of code
