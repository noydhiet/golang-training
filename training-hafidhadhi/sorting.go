package main

import "fmt"

func main() {

	urutanAcak := [6]int{10, 3, 2, 1, 8, 7}
	length := len(urutanAcak)
	for i := 1; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if urutanAcak[j] < urutanAcak[i] {
				tmp := urutanAcak[i]
				urutanAcak[i] = urutanAcak[j]
				urutanAcak[j] = tmp
			}
		}

	}
	fmt.Println(urutanAcak)

}
