package main

import (
	"fmt"
)

func main() {
	var i int
	isPrime := true
	fmt.Println("Masukan angka : ")

	fmt.Scanf("%d", &i)

	for j := 2; j <= i/2; j++ {
		if i%j == 0 {
			isPrime = false
		}
	}
	fmt.Println(isPrime)
	fmt.Println(i)

}
