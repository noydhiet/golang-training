package main

import "fmt"

func main() {
	var i int
	fmt.Println("Enter your value : ")

	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)
	}

	val := 0
	for j := 1; j <= i; j++ {
		if i%j == 0 {
			val++
		}
	}

	if val == 2 {
		fmt.Println("Bilangan prima")
	} else {
		fmt.Println("Bukan bilangan prima")
	}
}
