package main

import (
	"fmt"
)

// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter Your Name : ")
// 	name, _ := reader.ReadString("\n")
// 	fmt.Println(name)
// }

// func main() {

// 	var i int

// 	fmt.Println("Enter Number : ")

// 	fmt.Scanf("%d", &i)

// 	for j := 2; j <= i/2; j++ {
// 		if i%j == 0 {
// 			fmt.Println("Not")
// 			break
// 		}
// 	}

// 	fmt.Println("Prime")

// }

func main() {

	var i int

	fmt.Println("Enter Number : ")

	fmt.Scanf("%d", &i)
	k := 0
	for j := 1; j <= i; j++ {
		if i%j == 0 {
			k++
		}
	}
	if k == 2 {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}
