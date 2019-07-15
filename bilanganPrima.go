package main

import "fmt"

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	var value int

	fmt.Println("Enter year value : ")

	fmt.Scanf("%d", &value)

	fmt.Println(isPrime(value))
}
