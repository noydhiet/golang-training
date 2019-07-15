package main

import (
	"fmt"
)

func main() {

	var i int
	fmt.Println(" nomormu:")
	fmt.Scanf("%d", &i)
	a := 0
	for j := 1; j <= i; j++ {
		if i%j == 0 {
			a++
		}
	}
	if a == 2 {
		fmt.Println("Prima dong")
	} else {
		fmt.Println("bukan Prima")
	}
}
