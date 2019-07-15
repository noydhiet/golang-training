package main

import "fmt"

func main(){
	var a int = 10
	var b, c int
	a, b, c = multipleA(a, 10)
	fmt.Println(a, b, c)
	printA(a)
}

func multipleA(a int, b int) (result int, result2 int, result3 int){
	result = a*b
	return result, 10, 100
}

func printA(a int)  {
	fmt.Println(a)
}