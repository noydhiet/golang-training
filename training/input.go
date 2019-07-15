package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your name :")
	// name, _ := reader.ReadString('\n')
	// fmt.Println(name)

	// var i int
	// fmt.Println("Enter year : ")
	// fmt.Scanf("%d", &i)

	// fmt.Println(i)

	var i int
	// var n int
	fmt.Println("Enter angka : ")
	fmt.Scanf("%d", &i)

	// for n = 2 n< i n++

	if i%i == 0 && i%1 == i {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan")
	}

	// var i int
	// fmt.Println("enter year : ")

	// _, err := fmt.Scanf("%d", &i)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// k := 0

	// for j := 1; j <= i; j++ {
	// 	if i%j == 0 {
	// 		k++
	// 	}
	// }

	// if k == 2{
	// 	fmt.Println("Bilangan Prima")
	// } else {
	// 	fmt.Println("Bukan Bilangan Prima")
	// }
}
