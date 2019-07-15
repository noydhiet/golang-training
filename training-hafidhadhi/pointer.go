package main

import "fmt"

func main() {

	// address := new(string)
	// fmt.Println("Location variable: ", address)
	var a int = 10
	b := multipleAv2(a, 11)
	multipleA(&a, 10)
	printA(a)
	printA(b)
}

func multipleA(a *int, b int) {
	*a = *a * b
}

func printA(a int) {
	fmt.Println(a)
}

func multipleAv2(a int, b int) (result int) {
	result = a * b
	return result
}
