package main

import(
	"fmt"
)

func main(){

	var i int
	fmt.Println("Enter number : ")

	fmt.Scanf("%d", &i)
	
	for j :=2; j<=i/2; j++{
		if i%j == 0 {
			fmt.Println("Not")

		}else{
			fmt.Println("Primes")
		}
	}
}
