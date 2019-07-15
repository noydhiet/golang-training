package main

import "fmt"

//import "sort"

func main() {

	urutanAcak := [6]int{10, 3, 2, 1, 8, 7}
	temp := 0

	for i := 0; i <= 5; i++ {
		for j := i - 1; j <= 5; j++ {
			if urutanAcak[i] < urutanAcak[j] {
				temp = urutanAcak[i]
				urutanAcak[i] = urutanAcak[j]
				urutanAcak[j] = temp
			}
		}
	}

	for x := 0; x <= 5; x++ {
		fmt.Println("%d, ", urutanAcak[x])
	}
}
