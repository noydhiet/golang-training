package main
import "fmt"

func main() {
	
	// input a number
	fmt.Println("Enter year: ")

	fmt.Scanf("%d", &i)    // & itu untuk output dan input

	k := 0
	for j := 1; j <= i; j++ {
		if  i%j == 0 {
			k++
		}
	}
	
	if k == 2 {
		fmt.Println("Bilangan prima")
	}
	else {
		fmt.Println("Bilangan non prima")
	}

}

