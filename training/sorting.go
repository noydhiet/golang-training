package main

import "fmt"

func main() {

	urutanAcak := [6]int{10, 3, 2, 1, 8, 7}

	// for i := 0; i < len(urutanAcak); i++ {
	// 	fmt.Println(i, urutanAcak[i])

	for i, urutanAcak2 := range urutanAcak {

		fmt.Println(i, urutanAcak2)
	}

}
