package main

import "fmt"

func main() {
	urutanAcak := [6]int{10, 3, 2, 1, 8, 7}

	// sort.Ints(urutanAcak)
	// fmt.Println(urutanAcak)

	for i := 0; i < len(urutanAcak)-1; i++ {
		indexMin := i
		for j := i; j < len(urutanAcak); j++ {
			if urutanAcak[indexMin] > urutanAcak[j] {
				indexMin = j
			}
		}

		temp := urutanAcak[i]
		urutanAcak[i] = urutanAcak[indexMin]
		urutanAcak[indexMin] = temp
	}

	fmt.Println(urutanAcak)
}
