package main

import (
	"fmt"
)

func main() {

	urutanAcak := [6]int{10, 3, 2, 1, 8, 7}
	// sort.Ints(urutanAcak)
	// fmt.Println(urutanAcak)

	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < 6; i++ {
			if urutanAcak[i-1] > urutanAcak[i] {
				urutanAcak[i], urutanAcak[i-1] = urutanAcak[i-1], urutanAcak[i]
				swapped = true
			}
		}
	}
	fmt.Println(urutanAcak)

	// for i := 0; i < 6; i++ {
	// 	for j := i + 1; j < 6; j++ {
	// 		if urutanAcak[j] > urutanAcak[i] {
	// 			tmp := urutanAcak[i]
	// 			urutanAcak[i] = urutanAcak[j]
	// 			urutanAcak[j] = tmp
	// 		}
	// 	}
	// }
	// fmt.Println(urutanAcak)

}
