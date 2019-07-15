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

	if i/i == 1 && i/1 == i {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan")
	}

}
