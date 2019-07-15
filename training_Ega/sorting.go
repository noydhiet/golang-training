package main

import "fmt"

func main() {

	urutanAcak := [6]int{10, 3, 2, 1, 8, 7}

	// Manual
	end := 5

	for {
		if end == 0 {
			break
		}
		for i := 0; i < 5; i++ {
			if urutanAcak[i] > urutanAcak[i+1] {
				urutanAcak[i], urutanAcak[i+1] = urutanAcak[i+1], urutanAcak[i]
			}
		}
		end--
	}

	// Go method :
	// sort.Ints(urutanAcak)
	// fmt.Println(urutanAcak)

	fmt.Println(urutanAcak)
}
