package main

import "fmt"

// func multiplePointerA(a *int, b int) {
// 	*a = *a * b
// }

func main() {
	var a int = 10
	var b, c int
	address := new(string)
	fmt.Println("Location variable :", address)

	a, b, c = multipleA(a, 10)
	fmt.Println(a, b, c)

}

func multipleA(a int, b int) (result int, result2 int, result3 int) {
	result = a * b
	return result, 10, 100
}
