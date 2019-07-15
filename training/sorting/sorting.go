package main

import "fmt"

func main() {

	urutanAcak := [6]int{10, 3, 2, 1, 8, 7}

	end := len(urutanAcak) - 1

	for {

		if end == 0 {
			break
		}

		for i := 0; i < len(urutanAcak)-1; i++ {

			if urutanAcak[i] > urutanAcak[i+1] {
				urutanAcak[i], urutanAcak[i+1] = urutanAcak[i+1], urutanAcak[i]
			}

		}

		end -= 1

	}

	fmt.Println(urutanAcak)
}
