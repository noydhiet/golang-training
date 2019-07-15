package main

import "fmt"

func main() {
	//untuk memprint kata hello world
	//tidak dengan using parameter seperti input
	//fmt.Println("Hello world")

	//deklarasi variable dengan tipe string
	firstName, lastName := "hendy", "nurfrianto"
	address := new(string)
	//lastName := "nurfrianto"
	//print variable
	fmt.Printf("Hello %s %s !!", firstName, lastName)
	fmt.Printf(*address)
} // end of code
