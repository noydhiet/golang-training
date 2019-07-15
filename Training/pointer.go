package main

import "fmt"

func main() {
	var a int = 10
	var b int
	b= multipleA(a, 10)
	fmt.Println(a, b)
}
func multipleA(a int, b int) (result int) {
	result = a*b
	return result
}


